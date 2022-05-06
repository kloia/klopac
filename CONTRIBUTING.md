# Contributing to KloPaC

Thank you for your interest in contributing to KloPaC! 

KloPaC is MIT licensed and accepts contributions via GitHub pull requests. There are many areas we can use contributions - ranging from code, documentation, feature proposals, issue triage, samples, and content creation. This document outlines some conventions on development workflow and other resources to make it easier to get your contribution accepted.

We aim to build a vibrant and inclusive ecosystem. We want to make contributing to this project as easy and transparent as possible. 

The [Open Source Guides](https://opensource.guide/) website has resources for individuals, communities, and companies who want to learn how to run and contribute to an open source project. Contributors and people new to open source alike will find the following guides especially useful: 

- [How to Contribute to Open Source](https://opensource.guide/how-to-contribute/)
- [Building Welcoming Communities](https://opensource.guide/building-community/)

## We Use Github Flow, So All Code Changes Happen Through Pull Requests 

Pull requests are the best way to propose changes to the codebase.

### Prerequisites

1. Install Go

    KloPaC requires Go 1.18

### Use a Consistent Coding Style 

1. The coding style suggested by the Golang community is used in KloPaC. See the [style doc](https://github.com/golang/go/wiki/CodeReviewComments) for details.
1. Use two spaces for indentation rather than tabs in YAML files.
1. On the Dockerfile side, we mainly use Ubuntu 20.04 as the base image. We try not to use the root user and keep the number of layers and image size as low as possible.

## Getting Started 

1. [Fork](https://docs.github.com/en/get-started/quickstart/fork-a-repo#fork-an-example-repository) the repository. 

1. Play with the project, submit bugs, and submit patches!

### Contribution Flow 

1. Create a branch from where you want to base your work (usually main). 

1. Make your changes and arrange them in readable commits. 

1. Make sure your commit messages are in the proper format. 

1. Push your changes to the branch in your repository fork. 

1. Make sure all tests pass, and add any new tests as appropriate. 

1. Submit a pull request to the original repository. 

### Pull Request Flow 

The general flow for a pull request approval process is as follows: 

1. Author submits the pull request 

1. Reviewers and maintainers for the applicable code areas review the pull request and provide feedback that the author integrates 

1. Reviewers and/or maintainers signify their LGTM (Looks good to me) on the pull request 

1. A maintainer approves the pull request based on at least one LGTM from the previous step   

1. A maintainer merges the pull request into the target branch (main, release, etc.) 

### Report bugs using Github's issues 

We use GitHub issues to track public bugs. Report a bug by opening a new issue; it's that easy! 

Write bug reports with detail, background, and sample code. Great Bug Reports tend to have: 

- A quick summary and/or background 

- Steps to reproduce  

    - Be specific! 

    - Give a sample code if you can. 

- What you expected would happen 

- What actually happens 

- Notes (possibly including why you think this might be happening or stuff you tried that didn't work) 

## Code Quality

This project uses [pre-commits](https://pre-commit.com/) to ensure the quality of the code. Before committing, it's always a good idea to check the code for common programming mistakes, misspellings, and other potential errors. Every change is checked on CI, and it cannot be accepted if it does not pass the tests. Run:

```
pip3 intall pre-commit
pre-commit install
```

For more installation options visit the [pre-commits](https://pre-commit.com/).
