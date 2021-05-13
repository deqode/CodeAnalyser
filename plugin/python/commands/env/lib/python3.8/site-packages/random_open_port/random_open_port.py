import os
import sys
import argparse
import itertools
import requests
import random
from bs4 import BeautifulSoup

PORTS_OPTIONS = {
    "well-known" : ("Well-known ports",),
    "registered" : ("Registered ports",),
    "dynamic" : ("Dynamic, private or ephemeral ports",)
}

PORTS_OPTIONS['static'] = tuple(itertools.chain(PORTS_OPTIONS['well-known'], 
                                                PORTS_OPTIONS['registered']))

PORTS_OPTIONS['all'] = tuple(itertools.chain(PORTS_OPTIONS['static'],
                                             PORTS_OPTIONS['dynamic']))
PORTS_FILE = '.ports'
SCRIPT_PATH = os.path.dirname(os.path.abspath(__file__))
PORTS_FILE = os.path.join(SCRIPT_PATH, PORTS_FILE)

PORTS_BOUND = 2**15 + 2**14

def random_port(from_internet=False, which_ports='static'):
    """
    Get a random port that's not on the wikipedia list.
    
    Parameters
    ----------
    from_internet: bool, optional
        Whether to grab the list of unavaible ports from 
        Wikipedia or the already made cache. If True, this will 
        rescrape the ports from Wikipedia. Defaults to False.
    which_ports: str, optional
        If scraping from wikipedia, which ports to scrape.
        Defaults to 'static'. See get_wiki_ports for explanation.
    
    Returns
    -------
    int
        A random port not associated with anything in the wikipedia table.
    """
    ports = get_taken_ports(scrape=from_internet, which_ports=which_ports)
    while True:
        p = random.choice(range(PORTS_BOUND))
        if p not in ports:
            return p
    

def get_taken_ports(scrape=False, which_ports='static'):
    """
    Get the known ports either from the already written file,
    or from the web if desired or the file isn't populated.
    
    Parameters
    ----------
    scrape: bool, optional
        Whether to rescrape the ports list from Wikipedia. Defaults to
        False.
    which_ports: str, optional
        If rescraping, or not file has been written, which ports to scrape.
        Defaults to 'static'. See get_wiki_ports for explanantion of options.
        
    Returns
    -------
    tuple
    """
    if not _valid_ports_file() or scrape:
        ports = get_wiki_ports(which_ports)
        _write_ports_file(ports)
    else:
        ports = read_ports_file()
    return ports
        

def get_wiki_ports(which_ports='static'):
    """
    Get the known ports from the wikipedia table from this page
    "https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers"
    
    Parameters
    ----------
    which_ports: str, optional
        The options include 'well-known', 'registered', 'dynamic', 'static',
        and 'all'. 'well-known', 'registered', and 'dynamic' all correspond to
        their tables on Wikipedia. 'static' corresponds to both 'well-known'
        and 'registered', and 'all' corresponds to all 3. Defaults to 'static'.
    
    Returns
    -------
    tuple
        All of the ports from the desired tables.
    
    Raises
    ------
    ValueError: If which_ports is not a valid option.
    """
    if which_ports not in PORTS_OPTIONS:
        opts = _format_ports_options_str()
        raise ValueError("Expected either %s; got %s."%(opts, which_ports))
    
    url = "https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers"
    response = requests.get(url, timeout=5)
    soup = BeautifulSoup(response.content, 'html.parser')
    
    tables = soup.find_all('table', attrs={"class" : "wikitable"})
    tables = filter(lambda x : _get_caption(x) in PORTS_OPTIONS[which_ports], tables)
    scraped_ports = []
    for table in tables:
        scraped_ports.append(_get_port_numbers(table))
    return tuple(itertools.chain(*scraped_ports))

def read_ports_file():
    """
    Read the ports from the ports file.
    
    Returns
    -------
    tuple
    
    Raises
    ------
    OSError: If the ports file does not exist.
    """
    if not _valid_ports_file():
        raise OSError("The ports cache file is either empty or does not exist.")
    with open(PORTS_FILE, 'r') as ports_file:
        ports = ports_file.readlines()
    return tuple([int(port.strip()) for port in ports if port.strip()])

def _valid_ports_file():
    """
    Determine if the ports file exists and is valid.
    
    Returns
    -------
    bool
    """
    try:
        ports_file_size = os.path.getsize(PORTS_FILE)
        return ports_file_size > 0
    except OSError:
        return False

def _write_ports_file(ports):
    """
    Write the given ports to the ports file.
    
    Parameters
    ----------
    ports: iterable
    """
    with open(PORTS_FILE, 'w') as ports_file:
        ports_file.write('\n'.join([str(port) for port in ports]))

def _format_ports_options_str():
    """
    Format the list of options for the kwarg into a sentence format.
    """
    all_but_last = ', '.join(list(PORTS_OPTIONS.keys())[:-1])
    last = list(PORTS_OPTIONS.keys())[-1]
    return ', or '.join((all_but_last, last))
    
def _get_caption(table):
    """
    Get the caption of the associated wiki table.
    
    Parameters
    ----------
    table: BeautifulSoup
    
    Returns
    -------
    str
    """
    caption_text = None
    caption = table.find("caption")
    if caption:
        caption_text = caption.text.strip()
    return caption_text
    
def _get_port_numbers(table):
    """
    Get the port numbers from the html table on the Wikipedia page.
    
    Parameters
    ----------
    table: BeautfiulSoup
        This is the beautifulsoup representation of the html table
        associated with the ports.
    
    Returns
    -------
    list
        A list of the port numbers on that table.
    """
    ports = []
    for row in table.find_all('tr'):
        ports_col = row.find('td')
        if ports_col:
            ports_text = ports_col.text.strip()
            ports.extend(_parse_ports(ports_text))
    return ports

def _parse_ports(ports_text):
    """
    Handle the case where the entry represents a range of ports.
    
    Parameters
    ----------
    ports_text: str
        The text of the given port table entry.
        
    Returns
    -------
    tuple
        A tuple of all ports the text represents.
    """
    ports = ports_text.split('-')
    try:
        if len(ports) == 2:
            ports = tuple(range(int(ports[0]), int(ports[1]) + 1))
        else:
            ports = (int(ports[0]),)
    except ValueError:
        return ()
    
    return ports

def main():
    parser = argparse.ArgumentParser("Get a random port, not already associated with"\
                                     " any known programs via WikiPedia.")
    parser.add_argument("-i", "--internet", help="Use the current entries of Wikipedia to select an open port.",
                        action="store_true")
    
    parser.add_argument("-w", "--which", type=str, default='static', help="When getting the ports from Wikipedia, which ports should be considered. Options are %s. Defaults to static"%_format_ports_options_str())
    
    args = parser.parse_args()
    if args.which != 'static' and not args.internet:
        print("WARNING: You requested a different selection for ports, but you're not requesting"\
              " the ports from the internet. This value will be ignored unless passed with the -i flag.")
    
    try:
        port = random_port(args.internet, args.which)
        print("Random Port: %s"%port)
    except Exception as e:
        print("Exception Encountered")
        print(e)
        if hasattr(e, 'message'):
            print(e.message)
        sys.exit(1)
    else:
        sys.exit(0)

if __name__ == '__main__':
    main()
    
