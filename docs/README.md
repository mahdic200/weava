# Documentation for Weava

This is the source code for building `Weava` manual and you can make the newest version of docs for yourself locally or contribute to this project , which will be really appreciated ;) .

Let's get our hands on !

> [!NOTE]
> This README's instructions are compatible for distributions of Linux operating system e.g. Debian, Ubuntu, Mint, Fedora, MXLinux, Arch etc. , so If you are on Windows you might search for some command-lines on Google .

## Prerequisites

- python >= 3.12
- python3-venv (for linux)
- sphinx-build >= 5.3.0
- sphinx-autobuild >= 2024.10.03
- make (optional) >= 4.3

## Initialization

Open a terminal and navigate to this folder :

```bash
cd docs/
```

Then make a python virtual environment :

```bash
python3.12 -m venv venv
```

Then activate the `environment` :

```bash
source venv/bin/activate
```

## Dependency Installation

Now it's time to install our dependencies . we can do it in two ways which are below .

### Automated

```bash
(venv) user@hostname:~/path/to/docs
pip install -r installation.txt
```

### Manually

There is a list in `installation.txt` which you can copy paste it line by line and install it with pip .

But first upgrade the pip :

```bash
(venv) user@hostname:~/path/to/docs
$ pip install --upgrade pip
```

Then install `sphinx-autobuild` :

```bash
(venv) user@hostname:~/path/to/docs
pip install sphinx-autobuild==2024.10.3
```

Install `furo` :

```bash
(venv) user@hostname:~/path/to/docs
pip install furo==2024.8.6
```

## Development Server

Now that everything is ready you can start the development server by running the command :

```bash
make dev
```

Or you may wanna be the manual guy :) and run the development server manually :

```bash
(venv) user@hostname:~/path/to/docs
sphinx-autobuild source/ build/
```

After the server started without any errors , you can open your browser and enter this address [localhost:8000](http://localhost:8000) , your changes will be applied in realtime .

## Building Process

You can build the documentation using `make html` :

```bash
make html
```

The built documentation is in the `build` folder .
