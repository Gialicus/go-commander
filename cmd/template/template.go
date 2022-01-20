package template

import (
	"fmt"

	"giali.com/commander/cmd/util"
)

func PrintAlien(text string, value interface{}) {
	fmt.Printf("👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾\n")
	fmt.Printf("%s %s\n", text, value)
	fmt.Printf("👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾\n")
}
func PrintSkull(text string, value interface{}) {
	fmt.Printf("💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️\n")
	fmt.Printf("%s %s\n", text, value)
	fmt.Printf("💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️\n")
}
func PrintLogo(val string) {
	fmt.Printf("👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾\n")
	fmt.Printf("		  Hello, master %s\n", val)
	fmt.Printf("👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾\n")
	fmt.Printf("		     You are Welcome\n")
	fmt.Printf("👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾\n")
	fmt.Println("              #         #      #         #")
	fmt.Println("          ###          ##########          ###")
	fmt.Println("       #####           ##########           #####")
	fmt.Println("     #######           ##########           #######")
	fmt.Println("   ##########         ############         ##########")
	fmt.Println("  ####################################################")
	fmt.Println(" ######################################################")
	fmt.Println(" ######################################################")
	fmt.Println("########################################################")
	fmt.Println("########################################################")
	fmt.Println("########################################################")
	fmt.Println(" ######################################################")
	fmt.Println("  ########      ########################      ########")
	fmt.Println("   #######       #     #########      #       #######")
	fmt.Println("     ######             #######              ######")
	fmt.Println("       #####             #####              #####")
	fmt.Println("          ###             ###              ###")
	fmt.Println("            ##             #              ##")
	fmt.Printf("👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾\n")
	fmt.Printf("		   EvilCode Production\n")
	fmt.Printf("👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾\n")
}
func PrintMenuAndGetResult(labels []string) string {
	fmt.Printf("💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️\n")
	value := util.SelectQuestion("MENU", labels)
	fmt.Printf("💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️ 💀 ☠️\n")
	return value
}
