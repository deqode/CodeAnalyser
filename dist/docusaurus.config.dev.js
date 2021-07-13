"use strict";

var lightCodeTheme = require('prism-react-renderer/themes/github');

var darkCodeTheme = require('prism-react-renderer/themes/dracula');
/** @type {import('@docusaurus/types').DocusaurusConfig} */


module.exports = {
  title: 'Code Analyser',
  tagline: 'Complete Analysis of your code base ',
  url: 'https://deqode.github.io',
  baseUrl: '/CodeAnalyser/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/logo.png',
  organizationName: 'deqode',
  // Usually your GitHub org/user name.
  projectName: 'CodeAnalyser',
  // Usually your repo name.
  themeConfig: {
    navbar: {
      title: 'Code Analyser',
      logo: {
        alt: 'My Site Logo',
        src: 'img/logo.png'
      },
      items: [{
        type: 'doc',
        docId: 'intro',
        position: 'left',
        label: 'Documentation'
      }, {
        href: 'https://github.com/deqode/code-analyser',
        label: 'GitHub',
        position: 'right'
      }]
    },
    footer: {
      style: 'dark',
      links: [{
        title: 'Docs',
        items: [{
          label: 'Documentation',
          to: '/docs/intro'
        }]
      }, {
        title: 'Community',
        items: [{
          label: 'Stack Overflow',
          href: 'https://stackoverflow.com/questions/tagged/docusaurus'
        }, {
          label: 'Blog',
          href: 'https://deqode.com/blog'
        }, {
          label: 'Twitter',
          href: 'https://twitter.com/deqodesolutions'
        }]
      }, {
        title: 'More',
        items: [{
          label: 'GitHub',
          href: 'https://github.com/deqode'
        }, {
          label: 'Website',
          href: 'https://deqode.com/'
        }]
      }],
      copyright: "Copyright \xA9 ".concat(new Date().getFullYear(), " Deqode, Inc.")
    },
    prism: {
      theme: lightCodeTheme,
      darkTheme: darkCodeTheme
    }
  },
  presets: [['@docusaurus/preset-classic', {
    docs: {
      sidebarPath: require.resolve('./sidebars.js'),
      // Please change this to your repo.
      editUrl: 'https://github.com/deqode/code-analyser/edit/master/website/'
    },
    theme: {
      customCss: require.resolve('./src/css/custom.css')
    }
  }]]
};