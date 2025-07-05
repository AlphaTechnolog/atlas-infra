#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { AtlasInfraStack } from '../lib/atlas-infra-stack';

const app = new cdk.App();

new AtlasInfraStack(app, 'AtlasInfraStack', {
  stackName: 'AtlasInfraStack',
  description: 'General stack for the application',
  env: { account: process.env.CDK_DEFAULT_ACCOUNT, region: process.env.CDK_DEFAULT_REGION },
});