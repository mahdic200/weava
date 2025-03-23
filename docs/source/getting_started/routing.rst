Routing
#######

For routing you can simply use ``Routes/Routes.go`` file . open it and you can see the contents of the file is something like this :

.. code-block:: go

    func SetupRoutes(app *fiber.App) {
        app.Post("/login", AuthController.Login).Name("app.user.index")
        // adminGroup := app.Group("/admin", Auth.AuthMiddleware)
        adminGroup := app.Group("/admin")

        userGroup := adminGroup.Group("user")
        userGroup.Get("/", UserController.Index).Name("app.user.index")
        userGroup.Get("/show/:id", UserController.Show).Name("app.user.show")
        userGroup.Post("/store", UserValidation.Store(), UserController.Store).Name("app.user.store")

        app.Use("*", func(c *fiber.Ctx) error {
            return c.Status(404).JSON(fiber.Map{
                "message": "Not found !",
            })
        })
    }

Don't be afraid , there is nothing special about this piece of code , I will cover all of it for you in next chapters .

Defining Routes
---------------
Every route must be defined inside of the ``SetupRoutes`` function's scope with ``app`` variable . there are some examples of how you can do it . you can simply call one of the **Get** or **Post** methods on app variable and give it a route path and at least one and at last two handlers as second and third arguments respectively (I'll Cover this topic in other chapters that why just two handlers at last) . just like how I defined ``"/"`` route on ``adminGroup`` :

.. code-block:: go

    // another example on how you can define a route in root path without any Group
    app.Get("/", function(c *fiber.Ctx) {
        return c.Status(200).JSON(fiber.Map{
            "message": "hello world !",
        })
    })


.. important::
    Notice that there are also other methods like **Put**, **Patch**, **Delete** and **Head** that you can call , but for sake of simplicity we will just use **Get** and **Post** .
    The reason behind this is to avoid extra complexity and ultimately overwhelming with clarifying everything .

Grouping Routes
---------------

You can group your routes for more control, authentication, access-level system etc. first , you make your group with ``app`` variable , like ``admin`` which I wrote for you and put it in a variable and postfix it with word ``Group`` all in camel-case , you got it ? that's the rule for clean code in this repository , and please don't break it . like what I've done to ``adminGroup`` variable .

.. code-block:: go

    // snip ...
    adminGroup := app.Group("admin")
    // snip ...

Nested Groups
-------------

You can define nested route groups instead of passing route paths like this ``admin/user/index`` , you can write it much better and cleaner . to achieve this goal you first defined a group as I explained above , then use the group variable and then define another route group again , instead of using ``app`` variable , use the previous group you defined instead like the userGroup I defined :

.. code-block:: go

    // snip ...
    adminGroup := app.Group("admin")
    // snip ...
    userGroup := adminGroup.Group("user")
    // snip ...

You see ? life is much better in this way .
