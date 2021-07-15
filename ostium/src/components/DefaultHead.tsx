import Head from "next/head";
import React from "react";

export interface DefaultHeadProps {
  title?: string;
  description?: string;
  url?: string;
}

const defaults: DefaultHeadProps = {
  title: "LiveAssist",
  description:
    "A knowledge sharing platform made for programmers. Programmers can talk to each other with live websocket-based messaging and get the help they need.",
  url: "https://liveassist.satvikreddy.com/",
};

export const DefaultHead: React.FC<DefaultHeadProps> = ({
  title = defaults.title,
  description = defaults.description,
  url = defaults.url,
}) => {
  return (
    <div>
      <Head>
        <title>{title}</title>
        <meta name="title" content={title} />
        <meta name="description" content={description} />

        <meta property="og:type" content="website" />
        <meta property="og:url" content={url} />
        <meta property="og:title" content={title} />
        <meta property="og:description" content={description} />

        <meta property="twitter:url" content={url} />
        <meta property="twitter:title" content={title} />
        <meta property="twitter:description" content={description} />
      </Head>
    </div>
  );
};
