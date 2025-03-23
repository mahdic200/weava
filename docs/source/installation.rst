############
Installation
############

As you can see , the ``Weava`` is not a package , or module , it's a code template and architecture for you to make your job easier . in fact , there is nothing special about installing it , actually it's just downloading and placing it in a folder and initiating it .

Open Weava's `repository <https://github.com/mahdic200/weava>`_ and download the code . place it in somewhere you usually use for coding , like ``~/web-development`` folder .

Open the project with VSCode or any other editor you are comfortable with .
Copy the ``.env.example`` file and rename it to ``.env`` .
Generate a random key for ``JWT_SECRET`` variable , you can do it in anyways but I suggest ``openssl`` library :

.. code-block:: bash

    openssl rand -base64 24

it will generate a random string like this :

.. code-block::

    MlXM8+xbnXqgYT7SL44omX7ZAx6KrCdf

Start your Postgres service , in Linux it will be like this :

.. code-block:: bash
    
    sudo systemctl start postgresql

Copy it and put it in the ``JWT_SECRET=`` variable in ``.env`` file , here is an example :

.. code-block::

    # snip ...

    JWT_SECRET=MlXM8+xbnXqgYT7SL44omX7ZAx6KrCdf

    # snip ...

Put the username of your PostgreSQL user in ``DB_USER`` just like the ``JWT_SECRET`` variable :

.. code-block::

    # snip ...

    DB_USER=username

    # snip ...

Now put the password of your PostgreSQL user in ``DB_PASSWORD`` just like the ``DB_USER`` variable :


.. code-block::

    # snip ...

    DB_PASSWORD=somepassword

    # snip ...

Do the same for your PostgreSQL's port and host and database name in ``DB_PORT``, ``DB_HOST`` and ``DB_DBNAME`` variables respectively .

Congratulations ! you are done setting up your very first project with **Weava** !
