package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/redhat-developer/odo/pkg/application"
	"github.com/redhat-developer/odo/pkg/catalog"
	"github.com/redhat-developer/odo/pkg/component"
	"github.com/redhat-developer/odo/pkg/project"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	componentBinary string
	componentGit    string
	componentLocal  string
)

var componentCreateCmd = &cobra.Command{
	Use:   "create <component_type> [component_name] [flags]",
	Short: "Create a new component",
	Long: `Create a new component to deploy on OpenShift.

If component name is not provided, component type value will be used for the name.

A full list of component types that can be deployed is available using: 'odo component list'`,
	Example: `  # Create new Node.js component with the source in current directory. 
  odo create nodejs

  # Create new Node.js component named 'frontend' with the source in './frontend' directory
  odo create nodejs frontend --local ./frontend

  # Create new Node.js component with source from remote git repository.
  odo create nodejs --git https://github.com/openshift/nodejs-ex.git

  # Create new Wildfly component with binary named sample.war in './downloads' directory
  odo create wildfly wildly --binary ./downloads/sample.war

  # List of ready-to-use examples
  # for more examples, visit: https://github.com/redhat-developer/odo/blob/master/docs/examples.md
  odo create php --git https://github.com/openshift/cakephp-ex.git
  odo create python --git https://github.com/openshift/django-ex.git
	`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {

		stdout := color.Output
		log.Debugf("Component create called with args: %#v, flags: binary=%s, git=%s, local=%s", strings.Join(args, " "), componentBinary, componentGit, componentLocal)

		client := getOcClient()
		applicationName, err := application.GetCurrentOrGetCreateSetDefault(client)
		checkError(err, "")
		projectName := project.GetCurrent(client)

		checkFlag := 0

		if len(componentBinary) != 0 {
			checkFlag++
		}
		if len(componentGit) != 0 {
			checkFlag++
		}
		if len(componentLocal) != 0 {
			checkFlag++
		}

		if checkFlag > 1 {
			fmt.Println("The source can be either --binary or --local or --git")
			os.Exit(1)
		}

		//We don't have to check it anymore, Args check made sure that args has at least one item
		// and no more than two
		componentType := args[0]
		exists, err := catalog.Exists(client, componentType)
		checkError(err, "")
		if !exists {
			fmt.Printf("Invalid component type: %v\nRun 'odo catalog list' to see a list of supported components\n", componentType)
			os.Exit(1)
		}

		componentName := args[0]
		if len(args) == 2 {
			componentName = args[1]
		}
		//validate component name
		err = validateName(componentName)
		checkError(err, "")
		exists, err = component.Exists(client, componentName, applicationName, projectName)
		checkError(err, "")
		if exists {
			fmt.Printf("component with the name %s already exists in the current application\n", componentName)
			os.Exit(1)
		}

		if len(componentGit) != 0 {
			err := component.CreateFromGit(client, componentName, componentType, componentGit, applicationName)
			checkError(err, "")
			fmt.Printf("Component '%s' was created.\n", componentName)
			fmt.Printf("Triggering build from %s.\n\n", componentGit)
			err = component.Build(client, componentName, applicationName, true, true, stdout)
			checkError(err, "")
		} else if len(componentLocal) != 0 {
			// we want to use and save absolute path for component
			dir, err := filepath.Abs(componentLocal)
			checkError(err, "")
			fileInfo, err := os.Stat(dir)
			checkError(err, "")
			if !fileInfo.IsDir() {
				fmt.Println("Please provide a path to the directory")
				os.Exit(1)
			}
			err = component.CreateFromPath(client, componentName, componentType, dir, applicationName, "local")
			checkError(err, "")
			fmt.Printf("Please wait, creating %s component ...\n", componentName)
			err = component.Build(client, componentName, applicationName, false, true, stdout)
			checkError(err, "")
			fmt.Printf("Component '%s' was created.\n", componentName)
			fmt.Printf("To push source code to the component run 'odo push'\n")
		} else if len(componentBinary) != 0 {
			path, err := filepath.Abs(componentBinary)
			checkError(err, "")

			err = component.CreateFromPath(client, componentName, componentType, path, applicationName, "binary")
			checkError(err, "")
			fmt.Printf("Please wait, creating %s component ...\n", componentName)
			err = component.Build(client, componentName, applicationName, false, true, stdout)
			checkError(err, "")
			fmt.Printf("Component '%s' was created.\n", componentName)
			fmt.Printf("To push source code to the component run 'odo push'\n")
		} else {
			// we want to use and save absolute path for component
			dir, err := filepath.Abs("./")
			checkError(err, "")
			err = component.CreateFromPath(client, componentName, componentType, dir, applicationName, "local")
			checkError(err, "")
			fmt.Printf("Please wait, creating %s component ...\n", componentName)
			err = component.Build(client, componentName, applicationName, false, true, stdout)
			checkError(err, "")
			fmt.Printf("Component '%s' was created.\n", componentName)
			fmt.Printf("To push source code to the component run 'odo push'\n")
		}
		// after component is successfully created, set is as active
		err = component.SetCurrent(client, componentName, applicationName, projectName)
		checkError(err, "")
		fmt.Printf("\nComponent '%s' is now set as active component.\n", componentName)
	},
}

func init() {
	componentCreateCmd.Flags().StringVar(&componentBinary, "binary", "", "Binary artifact")
	componentCreateCmd.Flags().StringVar(&componentGit, "git", "", "Git source")
	componentCreateCmd.Flags().StringVar(&componentLocal, "local", "", "Use local directory as a source for component")

	// Add a defined annotation in order to appear in the help menu
	componentCreateCmd.Annotations = map[string]string{"command": "component"}
	componentCreateCmd.SetUsageTemplate(cmdUsageTemplate)

	rootCmd.AddCommand(componentCreateCmd)
}
