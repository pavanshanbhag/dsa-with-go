package main

import (
	"fmt"
	"os"
	"strings"

	arrayEasy "dsa-practice-problems/arrays-strings/easy"
	arrayMedium "dsa-practice-problems/arrays-strings/medium"
	backtrackmedium "dsa-practice-problems/backtracking/medium"
	dsEasy "dsa-practice-problems/data-structures/easy"
	dsMedium "dsa-practice-problems/data-structures/medium"
	dpeasy "dsa-practice-problems/dynamic-programming/easy"
	dpmedium "dsa-practice-problems/dynamic-programming/medium"
	graphEasy "dsa-practice-problems/graphs/easy"
)

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	switch os.Args[1] {
	case "arrays-easy":
		fmt.Println("🚀 Running Easy Array & String Problems")
		fmt.Println("=======================================")
		arrayEasy.DemonstrateEasyProblems()
		arrayEasy.ProblemComplexityAnalysis()

	case "arrays-medium":
		fmt.Println("🚀 Running Medium Array & String Problems")
		fmt.Println("=========================================")
		arrayMedium.DemonstrateMediumProblems()
		arrayMedium.ProblemComplexityAnalysis()

	case "ds-easy":
		fmt.Println("🚀 Running Easy Data Structure Problems")
		fmt.Println("=======================================")
		dsEasy.DemonstrateEasyDataStructures()
		dsEasy.ProblemComplexityAnalysis()

	case "ds-medium":
		fmt.Println("🚀 Running Medium Data Structure Problems")
		fmt.Println("=========================================")
		dsMedium.DemonstrateMediumDataStructures()
		dsMedium.ProblemComplexityAnalysis()

	case "arrays":
		fmt.Println("🚀 Running All Array & String Practice Problems")
		fmt.Println("===============================================")

		arrayEasy.DemonstrateEasyProblems()
		arrayEasy.ProblemComplexityAnalysis()

		fmt.Println("\n" + strings.Repeat("=", 60))

		arrayMedium.DemonstrateMediumProblems()
		arrayMedium.ProblemComplexityAnalysis()

	case "data-structures":
		fmt.Println("🚀 Running All Data Structure Practice Problems")
		fmt.Println("===============================================")

		dsEasy.DemonstrateEasyDataStructures()
		dsEasy.ProblemComplexityAnalysis()

		fmt.Println("\n" + strings.Repeat("=", 60))

		dsMedium.DemonstrateMediumDataStructures()
		dsMedium.ProblemComplexityAnalysis()

	case "graph-easy":
		fmt.Println("🚀 Running Easy Graph Problems")
		fmt.Println("==============================")
		graphEasy.DemonstrateEasyGraphs()
		graphEasy.ProblemComplexityAnalysis()

	case "graphs":
		fmt.Println("🚀 Running All Graph Practice Problems")
		fmt.Println("======================================")

		graphEasy.DemonstrateEasyGraphs()
		graphEasy.ProblemComplexityAnalysis()

	case "dp-easy":
		fmt.Println("🚀 Running Easy Dynamic Programming Problems")
		fmt.Println("============================================")
		dpeasy.DemonstrateEasyDP()
		dpeasy.ProblemComplexityAnalysis()

	case "dp-medium":
		fmt.Println("🚀 Running Medium Dynamic Programming Problems")
		fmt.Println("==============================================")
		dpmedium.DemonstrateMediumDP()
		dpmedium.ProblemComplexityAnalysis()

	case "dynamic-programming":
		fmt.Println("🚀 Running All Dynamic Programming Practice Problems")
		fmt.Println("====================================================")

		dpeasy.DemonstrateEasyDP()
		dpeasy.ProblemComplexityAnalysis()

		fmt.Println("\n" + strings.Repeat("=", 60))

		dpmedium.DemonstrateMediumDP()
		dpmedium.ProblemComplexityAnalysis()

	case "backtrack-medium":
		fmt.Println("🚀 Running Medium Backtracking Problems")
		fmt.Println("=======================================")
		backtrackmedium.DemonstrateBacktracking()
		backtrackmedium.ProblemComplexityAnalysis()

	case "backtracking":
		fmt.Println("🚀 Running All Backtracking Practice Problems")
		fmt.Println("==============================================")

		backtrackmedium.DemonstrateBacktracking()
		backtrackmedium.ProblemComplexityAnalysis()

	case "all":
		fmt.Println("🚀 Running All Practice Problems")
		fmt.Println("================================")

		fmt.Println("\n📚 ARRAYS & STRINGS")
		fmt.Println(strings.Repeat("-", 40))
		arrayEasy.DemonstrateEasyProblems()
		arrayEasy.ProblemComplexityAnalysis()

		fmt.Println("\n" + strings.Repeat("=", 60))

		arrayMedium.DemonstrateMediumProblems()
		arrayMedium.ProblemComplexityAnalysis()

		fmt.Println("\n📚 DATA STRUCTURES")
		fmt.Println(strings.Repeat("-", 40))
		dsEasy.DemonstrateEasyDataStructures()
		dsEasy.ProblemComplexityAnalysis()

		fmt.Println("\n" + strings.Repeat("=", 60))

		dsMedium.DemonstrateMediumDataStructures()
		dsMedium.ProblemComplexityAnalysis()

		fmt.Println("\n📚 GRAPHS")
		fmt.Println(strings.Repeat("-", 40))
		graphEasy.DemonstrateEasyGraphs()
		graphEasy.ProblemComplexityAnalysis()

		fmt.Println("\n📚 DYNAMIC PROGRAMMING")
		fmt.Println(strings.Repeat("-", 40))
		dpeasy.DemonstrateEasyDP()
		dpeasy.ProblemComplexityAnalysis()

		fmt.Println("\n" + strings.Repeat("=", 60))

		dpmedium.DemonstrateMediumDP()
		dpmedium.ProblemComplexityAnalysis()

		fmt.Println("\n📚 BACKTRACKING")
		fmt.Println(strings.Repeat("-", 40))
		backtrackmedium.DemonstrateBacktracking()
		backtrackmedium.ProblemComplexityAnalysis()

	case "help":
		showUsage()

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		showUsage()
	}
}

func showUsage() {
	fmt.Println("DSA Practice Problems - Complete Collection")
	fmt.Println("==========================================")
	fmt.Println()
	fmt.Println("Usage: go run main.go [command]")
	fmt.Println()
	fmt.Println("📚 Category Commands:")
	fmt.Println("  arrays         - Run all array & string problems")
	fmt.Println("  data-structures - Run all data structure problems")
	fmt.Println("  graphs         - Run all graph problems")
	fmt.Println("  dynamic-programming - Run all dynamic programming problems")
	fmt.Println("  backtracking   - Run all backtracking problems")
	fmt.Println("  all            - Run all problems across categories")
	fmt.Println()
	fmt.Println("🎯 Difficulty Commands:")
	fmt.Println("  arrays-easy    - Easy array & string problems (10 problems)")
	fmt.Println("  arrays-medium  - Medium array & string problems (10 problems)")
	fmt.Println("  ds-easy        - Easy data structure problems (10 problems)")
	fmt.Println("  ds-medium      - Medium data structure problems (10 problems)")
	fmt.Println("  graph-easy     - Easy graph problems (10 problems)")
	fmt.Println("  dp-easy        - Easy dynamic programming problems (10 problems)")
	fmt.Println("  dp-medium      - Medium dynamic programming problems (10 problems)")
	fmt.Println("  backtrack-medium - Medium backtracking problems (10 problems)")
	fmt.Println()
	fmt.Println("ℹ️  Help:")
	fmt.Println("  help           - Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run main.go arrays-easy")
	fmt.Println("  go run main.go ds-medium")
	fmt.Println("  go run main.go graph-easy")
	fmt.Println("  go run main.go arrays")
	fmt.Println("  go run main.go all")
	fmt.Println()
	fmt.Println("Testing:")
	fmt.Println("  go test ./arrays-strings/easy -v")
	fmt.Println("  go test ./arrays-strings/medium -v")
	fmt.Println("  go test ./data-structures/easy -v")
	fmt.Println("  go test ./data-structures/medium -v")
	fmt.Println("  go test ./graphs/easy -v")
	fmt.Println("  go test ./... -v")
	fmt.Println()
	fmt.Println("Benchmarking:")
	fmt.Println("  go test ./arrays-strings/easy -bench=.")
	fmt.Println("  go test ./data-structures/easy -bench=.")
	fmt.Println("  go test ./graphs/easy -bench=.")
}
