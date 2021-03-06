# Overview of the Odo (OpenShift Do) CLI Structure

> Example application

```sh
  # Creating and deploying a Node.js project
  git clone https://github.com/openshift/nodejs-ex && cd nodejs-ex
  odo create nodejs
  odo push

  # Accessing your Node.js component
  odo url create
``` 

Odo (OpenShift Do) is a CLI tool for running OpenShift applications in a fast and automated matter. Odo reduces the complexity of deployment by adding iterative development without the worry of deploying your source code.

Find more information at https://github.com/redhat-developer/odo

# Syntax

#### List of Commands

|           NAME            |                  DESCRIPTION                   |
|---------------------------|------------------------------------------------|
| [app](#app)               | Perform application operations                 |
| [catalog](#catalog)       | Catalog related operations                     |
| [completion](#completion) | Output shell completion code                   |
| [component](#component)   | Components of application.                     |
| [create](#create)         | Create a new component                         |
| [delete](#delete)         | Delete an existing component                   |
| [describe](#describe)     | Describe the given component                   |
| [link](#link)             | Link target component to source component      |
| [list](#list)             | List all components in the current application |
| [log](#log)               | Retrieve the log for the given component.      |
| [project](#project)       | Perform project operations                     |
| [push](#push)             | Push source code to a component                |
| [storage](#storage)       | Perform storage operations                     |
| [update](#update)         | Update the source code path of a component     |
| [url](#url)               | Expose component to the outside world          |
| [version](#version)       | Print the client version information           |
| [watch](#watch)           | Watch for changes, update component on change  |


#### CLI Structure

```sh
odo --verbose : Odo (Openshift Do)
    app --short : Perform application operations
        create : Create an application
        delete --force : Delete the given application
        describe : Describe the given application
        get --short : Get the active application
        list : Lists all the applications
        set : Set application as active
    catalog : Catalog related operations
        list : List all available component types.
        search : Search component type in catalog
    completion : Output shell completion code
    component --short : Components of application.
        get --short : Get currently active component
        set : Set active component.
    create --binary --git --local : Create a new component
    delete --force : Delete an existing component
    describe : Describe the given component
    link --component : Link target component to source component
    list : List all components in the current application
    log : Retrieve the log for the given component.
    project --short : Perform project operations
        create : Create a new project
        get --short : Get the active project
        list : List all the projects
        set --short : Set the current active project
    push --local : Push source code to a component
    storage : Perform storage operations
        create --component --path --size : Create storage and mount to a component
        delete --force : Delete storage from component
        list --component : List storage attached to a component
        mount --component --path : mount storage to a component
        unmount --component : Unmount storage from the current component
    update --binary --git --local : Update the source code path of a component
    url : Expose component to the outside world
        create : Create a URL for a component
        delete --force : Delete a URL
        list --application --component : List URLs
    version : Print the client version information
    watch : Watch for changes, update component on change

```

## app

`app`

> Example using app

```sh
  # Create an application
  odo app create myapp
	
  # Get the currently active application
  odo app get
	
  # Delete the application
  odo app delete myapp
	
  # Describe webapp application,
  odo app describe webapp
	
  # List all applications
  odo app list
	
  # Set an application as active
  odo app set myapp
	
```


Performs application operations related to your OpenShift project.

## catalog

`catalog [options]`

> Example using catalog

```sh
  # Get the supported components
  odo catalog list

  # Search for a supported component
  odo search nodejs

  # Search for a component
  odo catalog search pyt
	
```


Catalog related operations

## completion

`completion SHELL`

> Example using completion

```sh
  # Bash autocompletion support
  source <(odo completion bash)

  # Zsh autocompletion support
  source <(odo completion zsh)

```


Generates shell completion code.

Auto completion supports both bash and zsh. Output is to STDOUT.

## component

`component`

> Example using component

```sh
  # Get the currently active component
  odo component get
	
  # Set component named 'frontend' as active
  odo set component frontend
  
```




## create

`create <component_type> [component_name] [flags]`

> Example using create

```sh
  # Create new Node.js component with the source in current directory. 
  odo create nodejs

  # Create new Node.js component named 'frontend' with the source in './frontend' directory
  odo create nodejs frontend --local ./frontend

  # Create new Node.js component with source from remote git repository.
  odo create nodejs --git https://github.com/openshift/nodejs-ex.git

  # Create new Wildfly component with binary named sample.war in './downloads' directory
  odo create wildfly wildly --binary ./downloads/sample.war

  # Create a Ruby component
  odo create ruby
	
  # Create a Python component
  odo create python
	
```


Create a new component to deploy on OpenShift.

If component name is not provided, component type value will be used for the name.

A full list of component types that can be deployed is available using: 'odo component list'

## delete

`delete <component_name>`

> Example using delete

```sh
  # Delete component named 'frontend'. 
  odo delete frontend
	
```


Delete an existing component.

## describe

`describe [component_name]`

> Example using describe

```sh
  # Describe nodejs component,
  odo describe nodejs
	
```


Describe the given component.

## link

`link <target component> --component [source component]`

> Example using link

```sh
  # Link current component to a component 'mariadb'
  odo link mariadb

  # Link 'mariadb' component to 'nodejs' component
  odo link mariadb --component nodejs
	
```


Link target component to source component

If source component is not provided, the link is created to the current active
component.

In the linking process, the environment variables containing the connection
information from target component are injected into the source component and
printed to STDOUT.


## list

`list`

> Example using list

```sh
  # List all components in the application
  odo list
	
```


List all components in the current application.

## log

`log [component_name]`

> Example using log

```sh
  # Get the logs for the nodejs component
  odo log nodejs
	
```


Retrieve the log for the given component.

## project

`project [options]`

> Example using project

```sh
  # Set the current active project
  odo project set myproject
	
  # Create a new project
  odo project create myproject
	
  # List all the projects
  odo project list
	
  # Get the active project
  odo project get
	
```


Perform project operations

## push

`push [component name]`

> Example using push

```sh
  # Push source code to the current component
  odo push

  # Push data to the current component from the original source.
  odo push

  # Push source code in ~/mycode to component called my-component
  odo push my-component --local ~/mycode
	
```


Push source code to a component.

## storage

`storage`

> Example using storage

```sh
  # Create storage of size 1Gb to a component
  odo storage create mystorage --path=/opt/app-root/src/storage/ --size=1Gi
	
  # Delete storage mystorage from the currently active component
  odo storage delete mystorage

  # Delete storage mystorage from component 'mongodb'
  odo storage delete mystorage --component mongodb

  # Unmount storage 'dbstorage' from current component
  odo storage unmount dbstorage

  # Unmount storage 'database' from component 'mongodb'
  odo storage unmount database --component mongodb

  # List all storage attached or mounted to the current component and 
  # all unattached or unmounted storage in the current application
  odo storage list
	
```


Perform storage operations

## update

`update`

> Example using update

```sh
  # Change the source code path of a currently active component to local (use the current directory as a source)
  odo update --local

  # Change the source code path of the frontend component to local with source in ./frontend directory
  odo update frontend --local ./frontend

  # Change the source code path of a currently active component to git 
  odo update --git https://github.com/openshift/nodejs-ex.git

  # Change the source code path of the component named node-ex to git
  odo update node-ex --git https://github.com/openshift/nodejs-ex.git

  # Change the source code path of the component named wildfly to a binary named sample.war in ./downloads directory
  odo update wildfly --binary ./downloads/sample.war
	
```


Update the source code path of a component

## url

`url`

> Example using url

```sh
  # Create a URL for the current component.
  odo url create

  # Create a URL for a specific component
  odo url create mycomponent
	
  # Delete a URL to a component
  odo url delete myurl
	
 # List the available URLs
  odo url list
	
```


Expose component to the outside world.

The URLs that are generated using this command, can be used to access the deployed components from outside the cluster.

## version

`version`

> Example using version

```sh
  # Print the client version of Odo
  odo version
	
```


Print the client version information

## watch

`watch [component name]`

> Example using watch

```sh
  # Watch for changes in directory for current component
  odo watch

  # Watch for changes in directory for component called frontend 
  odo watch frontend
	
```


Watch for changes, update component on change.


