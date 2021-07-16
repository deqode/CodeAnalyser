const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').DocusaurusConfig} */
module.exports = {
  title: 'Code Analyser',
  tagline: 'A tool for complete analysis of your code & its dependencies ',
  url: 'https://deqode.github.io',
  baseUrl: '/CodeAnalyser/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/logo.png',
  organizationName: 'deqode', // Usually your GitHub org/user name.
  projectName: 'CodeAnalyser', // Usually your repo name.
  themeConfig: {
    navbar: {
      title: 'Code Analyser',
      logo: {
        alt: 'Code Analyser',
        src: 'img/logo.png',
      },
      items: [
        {
          type: 'doc',
          docId: 'intro',
          position: 'left',
          label: 'Documentation',
        },
        {
          href: 'https://github.com/deqode/code-analyser',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Documentation',
              to: '/docs/intro',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Stack Overflow',
              href: 'https://stackoverflow.com/questions/tagged/docusaurus',
            },
            {
              label: 'Blog',
              href: 'https://deqode.com/blog',
            },
            {
              label: 'Twitter',
              href: 'https://twitter.com/deqodesolutions',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'GitHub',
              href: 'https://github.com/deqode',
            },
            {
              label: 'Website',
              href: 'https://deqode.com/',
            },
          ],
        },

      ],
      logo: {
        alt: 'Deqode logo',
        src: 'img/deq.svg',
        href: 'https://deqode.com/',
      },
      copyright: `Copyright © ${new Date().getFullYear()} Deqode, Inc.`,
    },
    prism: {
      theme: lightCodeTheme,
      darkTheme: darkCodeTheme,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl:
            'https://github.com/deqode/CodeAnalyser/tree/documentation',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
};
