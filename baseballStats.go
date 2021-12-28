/*
 * Course: CS 424-01
 * Assignment: Program 1
 * Purpose: This program reads baseball player data from a file and computes and displays statistics for each player.
 * Tested with go version go1.17.1 windows/amd64
 * Author: Matthew Hise
 * Email: mrh0036@uah.edu
 * Date Created: 09/30/21
 * Date Modified: 10/02/21
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// PlayerRecord A structure for holding all baseball player data
type PlayerRecord struct {
	fName string		// The player's first name
	lName string		// The player's last name
	atPlates int		// The number of plate appearances for this player
	atBats int			// The number of at bats for the player
	singles int			// The number of singles for this player
	doubles int			// The number of doubles for this player
	triples int			// The number of triples for this player
	homeRuns int		// The number of home runs for this player
	walks int			// The number of walks for this player
	hitByPitches int	// The number of hit-by-pitches for this player
}

/* Computes a player's batting average
 *
 * Parameters:
 * record - The PlayerRecord to compute the batting average for
 *
 * Returns:
 * float64 - the player's batting average
 */
func computeBattingAvg(record PlayerRecord) float64 {
	return float64(record.singles + record.doubles + record.triples + record.homeRuns) / float64(record.atBats)
}

/* Computes a player's slugging percentage
 *
 * Parameters:
 * record - The PlayerRecord to compute the slugging percentage for
 *
 * Returns:
 * float64 - the player's slugging percentage
 */
func computeSluggingPct(record PlayerRecord) float64 {
	return float64(record.singles + 2 * record.doubles + 3 * record.triples + 4 * record.homeRuns) / float64(record.atBats)
}

/* Computes a player's on-base percentage
 *
 * Parameters:
 * record - The PlayerRecord to compute the on-base percentage for
 *
 * Returns:
 * float64 - the player's on-base percentage
 */
func computeOnbasePct(record PlayerRecord) float64 {
	return float64(record.singles + record.doubles + record.triples + record.homeRuns + record.walks + record.hitByPitches) / float64(record.atPlates)
}

/* Computes a player's on-base plus slugging
 *
 * Parameters:
 * record - The PlayerRecord to compute the on-base plus slugging for
 *
 * Returns:
 * float64 - the player's on-base plus slugging
 */
func computeOnbaseSlugging(record PlayerRecord) float64 {
	return computeOnbasePct(record) + computeSluggingPct(record) // Compute and add the two percentages forming this stat
}

// ByOPS implements sort.Interface based on each player's On-base Plus Slugging
type ByOPS []*PlayerRecord

// Implement all needed functions for sorting
func (o ByOPS) Len() int           { return len(o) }
func (o ByOPS) Less(i, j int) bool { return computeOnbaseSlugging(*o[i]) > computeOnbaseSlugging(*o[j]) }
func (o ByOPS) Swap(i, j int)      { *o[i], *o[j] = *o[j], *o[i] }


// Simple function to handle errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Main function
func main() {
	// Print the welcome instructions and prompt for the user to enter an input filename
	fmt.Println("Welcome to the player statistics calculator program. This program will read players from an input data file.\n" +
		"You will provide the name of your input file, and the program will store all of the players in a list, compute\n" +
		"each player's averages, and write the resulting team report to the screen.\n\nEnter the name of your input file: ")

	var records []*PlayerRecord			// A slice of pointers to player records
	var playerRecord *PlayerRecord		// A pointer to a player record

	var filename string					// The name of the file to read
	fmt.Scanf("%s", &filename)	// Get the filename from the user

	// Open the file and handle any errors
	file, err := os.Open(filename)
	check(err)
	defer file.Close()					// Ensure the file will be closed

	scanner := bufio.NewScanner(file)	// Create a new scanner for the file

	// For each line of the player data file, store that data in a new player object
	for scanner.Scan() {
		playerRecord = new(PlayerRecord)					// Create a new player record and point to it
		dataElements := strings.Fields(scanner.Text())		// Split the current line on whitespace

		// Store the data from the line into the new player object, converting data types as necessary
		playerRecord.fName = dataElements[0]
		playerRecord.lName = dataElements[1]
		playerRecord.atPlates, _ = strconv.Atoi(dataElements[2])
		playerRecord.atBats, _ = strconv.Atoi(dataElements[3])
		playerRecord.singles, _ = strconv.Atoi(dataElements[4])
		playerRecord.doubles, _ = strconv.Atoi(dataElements[5])
		playerRecord.triples, _ = strconv.Atoi(dataElements[6])
		playerRecord.homeRuns, _ = strconv.Atoi(dataElements[7])
		playerRecord.walks, _ = strconv.Atoi(dataElements[8])
		playerRecord.hitByPitches, _ = strconv.Atoi(dataElements[9])

		// Add the new record to the slice
		records = append(records, playerRecord)
	}

	// Sort the players by on-base plus slugging, highest to lowest
	sort.Sort(ByOPS(records))

	// Print the data report summary header
	fmt.Println("\nBASEBALL TEAM REPORT --- " + strconv.Itoa(len(records)) + " PLAYERS FOUND IN FILE\n")
	fmt.Printf("%-20s : %10s %10s %10s %8s\n", "     PLAYER NAME", "AVERAGE", "SLUGGING", "ONBASE%", "OPS")
	fmt.Printf("------------------------------------------------------------------\n")

	// Print the statistics for each player sorted from highest to lowest OPS
	for _, record := range records {
		fmt.Printf("%20s : %9.3f %10.3f %10.3f %10.3f\n", record.lName + ", " + record.fName, computeBattingAvg(*record),
			computeSluggingPct(*record), computeOnbasePct(*record), computeOnbaseSlugging(*record))
	}

}
