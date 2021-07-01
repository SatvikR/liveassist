import { Heading } from "@chakra-ui/layout";
import React, { useEffect } from "react";
import { api } from "../lib/api";

export interface ChannelsProps {}

export const Channels: React.FC<ChannelsProps> = ({}) => {
  useEffect(() => {
    api.channels.list().then((data) => console.log(data));
  }, []);

  return <Heading>Active Channels</Heading>;
};
