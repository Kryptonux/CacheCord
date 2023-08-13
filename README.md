<!-- Improved compatibility of back to top link: See: https://github.com/Kryptonux/CacheCord/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/Kryptonux/CacheCord">
    <img src="images/logo.png" alt="Logo" width="160" height="160">
  </a>

  <h3 align="center">CacheCord</h3>

  <p align="center">
     A simple Go-utility for Linux users to recover cached files from Discord
    <br />
    <a href="https://github.com/Kryptonux/CacheCord"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    ·
    <a href="https://github.com/Kryptonux/CacheCord/issues">Report Bug</a>
    ·
    <a href="https://github.com/Kryptonux/CacheCord/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
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
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![CacheCord][product-screenshot]](https://github.com/Kryptonux/CacheCord)

CacheCord represents a straightforward initiative aimed at retrieving and reinstating cached files from Discord. It is essential to emphasize that this tool is solely intended for testing purposes, specifically designed to assess Discord's cache and storage mechanisms.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple example steps.

### Prerequisites

* Go-Lang 

### Installation

1: Clone the repo
   ```sh
   git clone https://github.com/Kryptonux/CacheCord.git
   ```
2: Install GO packages
   ```sh
   go mod tidy
   ```
3: Build CacheCord
   ```sh
   go build .
   ```
4.1: Start dumping cache with CLI
   ```sh
   ./cachecord -f png,mp4 -o "/your/output/path"
   ```
4.2: Start dumping cache with a interactive prompt
   ```sh
   ./cachecord -i
   ```
   
<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Email: contact@uwuu.nu
Project Link: [https://github.com/Kryptonux/CacheCord](https://github.com/Kryptonux/CacheCord)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/Kryptonux/CacheCord.svg?style=for-the-badge
[contributors-url]: https://github.com/Kryptonux/CacheCord/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/Kryptonux/CacheCord.svg?style=for-the-badge
[forks-url]: https://github.com/Kryptonux/CacheCord/network/members
[stars-shield]: https://img.shields.io/github/stars/Kryptonux/CacheCord.svg?style=for-the-badge
[stars-url]: https://github.com/Kryptonux/CacheCord/stargazers
[issues-shield]: https://img.shields.io/github/issues/Kryptonux/CacheCord.svg?style=for-the-badge
[issues-url]: https://github.com/Kryptonux/CacheCord/issues
[license-shield]: https://img.shields.io/github/license/Kryptonux/CacheCord.svg?style=for-the-badge
[license-url]: https://github.com/Kryptonux/CacheCord/blob/master/LICENSE.txt
[product-screenshot]: images/screenshot.png
