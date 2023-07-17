<!--
  Created by: Ethan Baker (contact@ethanbaker.dev)
  
  Adapted from:
    https://github.com/othneildrew/Best-README-Template/

Here are different preset "variables" that you can search and replace in this template.
`project_title`
`project_description`
`documentation_link`
-->

<div id="top"></div>


<!-- PROJECT SHIELDS/BUTTONS -->
![1.0.0](https://img.shields.io/badge/status-1.0.0-red)
[![GoDoc](https://godoc.org/github.com/ethanbaker/epick?status.svg)](https://godoc.org/github.com/ethanbaker/epick)
[![Go Report Card](https://goreportcard.com/badge/github.com/ethanbaker/epick)](https://goreportcard.com/report/github.com/ethanbaker/epick)
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]


<!-- PROJECT SPECIFIC BUTTONS -->
<!-- Netlify buttons:
[![Netlify Status]()]()
-->

<!-- Golang specific buttons -->

<!-- PROJECT LOGO -->
<br><br><br>
<div align="center">
  <!-- 
  <a href="https://github.com/ethanbaker/epick/docs/logo.png">
    <img src="" alt="Logo" width="80" height="80">
  </a>
  -->

  <h3 align="center">Epick</h3>

  <p align="center">
    A simple emoji picker for the terminal
  </p>
</div>


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>


<!-- ABOUT -->
## About

![Project demonstration image][product-screenshot]

Epick was designed to reduce the hassle of having to search up the copy and paste for different emojis whenever you want an emoji in a text document. For instance, adding emojis to blog posts in the terminal can be so much easier when using Epick, as you no longer need to copy and paste a character after searching for the perfect emoji for 5 minutes. All you need to do is run the `epick` command, find the preferred emoji, and continue coding!

<p align="right">(<a href="#top">back to top</a>)</p>


### Built With

* [Golang](https://go.dev/learn/)
* [Cview](https://code.rocketnine.space/tslocum/cview)

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

In order to download and start using epick as a terminal command, you must first clone the git repository to your local system.

In order to use epick in a different application, import it by using the import statement:
`import "github.com/ethanbaker/epick"`


### Prerequisites

* Go is installed
* Your terminal is able to render emojis

### Installation

1. Clone the repository (`git clone git@github.com:ethanbaker/epick.git`)
1. Navigate into the cloned directory
1. Run `go install`

Epick should now be installed!

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- USAGE EXAMPLES -->
## Usage

Epick contains multiple pages of different emojis, each of which represents a unique emoji category. These categories are:

* Smileys and Emotion
* People and Body
* Animals and Nature
* Food and Drink
* Travel and Places
* Activities
* Objects
* Symbols
* Flags

You can navigate between different emojis using standard vim bindings (j, j, k, l, g, G).

You can switch between emoji categories by pressing 'C' to go forwards and 'c' to go backwards.

You can quit the app by pressing 'q' or Escape.

You can search for emojis by pressing '?'. This brings up a menu where emojis can be searched for with autocomplete. If you enter a partially incomplete phrase, you can navigate between matching emojis by pressing 'n' to go forwards and 'N' to go backwards.


<p align="right">(<a href="#top">back to top</a>)</p>


<!-- ROADMAP -->
## Roadmap

- [ ] Documentation
- [ ] Quality-ensurance Refactor
- [ ] Testing

See the [open issues][issues-url] for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONTRIBUTING -->
## Contributing

For issues and suggestions, please include as much useful information as possible.
Review the [documentation][documentation-url] and make sure the issue is actually
present or the suggestion is not included. Please share issues/suggestions on the
[issue tracker][issues-url].

For patches and feature additions, please submit them as [pull requests][pulls-url]. 
Please adhere to the [conventional commits][conventional-commits-url]. standard for
commit messaging. In addition, please try to name your git branch according to your
new patch. [These standards][conventional-branches-url] are a great guide you can follow.

You can follow these steps below to create a pull request:

1. Fork the Project
2. Create your Feature Branch (`git checkout -b branch_name`)
3. Commit your Changes (`git commit -m 'commit_message'`)
4. Push to the Branch (`git push origin branch_name`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- LICENSE -->
## License

This project uses the Apache 2.0 license.

You can find more information in the `LICENSE` file.

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Ethan Baker - contact@ethanbaker.dev - [LinkedIn][linkedin-url]

Project Link: [https://github.com/ethanbaker/epick][project-url]

<p align="right">(<a href="#top">back to top</a>)</p>


<p align="right">(<a href="#top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/ethanbaker/epick.svg
[forks-shield]: https://img.shields.io/github/forks/ethanbaker/epick.svg
[stars-shield]: https://img.shields.io/github/stars/ethanbaker/epick.svg
[issues-shield]: https://img.shields.io/github/issues/ethanbaker/epick.svg
[license-shield]: https://img.shields.io/github/license/ethanbaker/epick.svg
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?logo=linkedin&colorB=555

[contributors-url]: <https://github.com/ethanbaker/epick/graphs/contributors>
[forks-url]: <https://github.com/ethanbaker/epick/network/members>
[stars-url]: <https://github.com/ethanbaker/epick/stargazers>
[issues-url]: <https://github.com/ethanbaker/epick/issues>
[pulls-url]: <https://github.com/ethanbaker/epick/pulls>
[license-url]: <https://github.com/ethanbaker/epick/blob/master/LICENSE>
[linkedin-url]: <https://linkedin.com/in/ethandbaker>
[project-url]: <https://github.com/ethanbaker/epick>

[product-screenshot]: ./docs/demonstration.png
[documentation-url]: <https://documentation_link>

[conventional-commits-url]: <https://www.conventionalcommits.org/en/v1.0.0/#summary>
[conventional-branches-url]: <https://docs.microsoft.com/en-us/azure/devops/repos/git/git-branching-guidance?view=azure-devops>