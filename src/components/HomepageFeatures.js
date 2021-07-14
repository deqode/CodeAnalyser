import React from 'react';
import clsx from 'clsx';
import styles from './HomepageFeatures.module.css';

const FeatureList = [
  {
    title: 'Plugin Architecture',
    Svg: require('../../static/img/plug-in.svg').default,
    description: (
      <>
        Code Analyser lets you focus on your detection, and we&apos;ll do the chores. Go
        ahead and make your own  <code>plug-in</code> .
      </>
    ),
  },
  {
    title: 'Multilingual',
    Svg: require('../../static/img/multilingual.svg').default,
    description: (
      <>
        Code Analyser can be extended in any language by just ading new plugin.
      </>
    ),
  },
  {
    title: 'Powered by Deqode',
    Svg: require('../../static/img/deq.svg').default,
    description: (
      <>Constanly updated and maintained by developers at Deqode.  </>
    ),
  },
];

function Feature({Svg, title, description}) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} alt={title} />
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
