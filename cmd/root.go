package cmd

var rootCmd = &cobra.Command{
	Use:   "gkp",
	Short: "gkp is a tool for manipulating keepass kpx files",
	Long: `Rex is learning Go and writing a KeePass tool. 
	       gkp can be used to read, write and store files in Keepass kpx files.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
