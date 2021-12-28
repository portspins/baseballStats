#=
program2hise:
- Julia version: 1.6.3
- Tested on: Windows 10 Pro
- Course: CS 424-01
- Assignment: Program 2
- Purpose: This program reads baseball player data from a file and computes and displays statistics for each player.
- Author: Matthew Hise
- Email: mrh0036@uah.edu
- Date Created: 2021-10-18
- Date Modified: 2021-10-30
=#

using Printf

struct PlayerRecord
	fName::String			# The player's first name
	lName::String			# The player's last name
	atPlates::UInt32		# The player's number of plate appearances
	atBats::UInt32			# The player's number of times at bat
	singles::UInt32			# The player's number of singles
	doubles::UInt32			# The player's number of doubles
	triples::UInt32			# The player's number of triples
	homeRuns::UInt32		# The player's number of home runs
	walks::UInt32			# The player's number of walks
	hitByPitches::UInt32	# The player's number of hit by pitches
end

# Calculate the batting average for a player
calcBattingAvg(p::PlayerRecord) = (p.singles + p.doubles + p.triples + p.homeRuns) / p.atBats

# Calculate the slugging percentage for a player
calcSluggingPct(p::PlayerRecord) = (p.singles + 2 * p.doubles + 3 * p.triples + 4 * p.homeRuns) / p.atBats

# Calculate the on-base percentage for a player
calcOBP(p::PlayerRecord) = (p.singles + p.doubles + p.triples + p.homeRuns + p.walks + p.hitByPitches) / p.atPlates

# Calculate the on-base plus slugging for a player
calcOPS(p::PlayerRecord) = calcSluggingPct(p) + calcOBP(p)

# Print out stats for all players in a list of players
# Parameters - A list of playerRecords, a string defining the attribute or method the list is sorted by
function printAllPlayerStats(players, sortedBy::String = "custom ordering")
	println("\nBASEBALL TEAM REPORT (SORTED BY ", uppercase(sortedBy), ")\n")
	@printf("%-20s : %10s %10s %10s %8s\n", "     PLAYER NAME", "AVERAGE", "SLUGGING", "ONBASE%", "OPS")
	println("------------------------------------------------------------------")

	for player in players
		@printf("%20s : %9.3f %10.3f %10.3f %10.3f\n", string(player.lName, ", ", player.fName),
		calcBattingAvg(player), calcSluggingPct(player), calcOBP(player), calcOPS(player))
	end
end

# Print introductory message
println("""Welcome to the player statistics calculator program. This program will read players from an input data file.
		You will provide the name of your input file, and the program will store all of the players in a list, compute
		each player's averages, and write the resulting team report to the screen.\n\nEnter the name of your input file:""")

# Get the file path and initialize the list for player records
filename = readline()
playerRecords = PlayerRecord[]

# Open the file and read in each player's data
open(filename) do f
	# For each line in the file, read in the data to a player record and add it to the list
	while ! eof(f)
		s = readline(f)
		playerData = split(s, " ")
		newPlayer = PlayerRecord(playerData[1], playerData[2], parse(UInt32, playerData[3]),
		parse(UInt32, playerData[4]), parse(UInt32, playerData[5]), parse(UInt32, playerData[6]),
		parse(UInt32, playerData[7]), parse(UInt32, playerData[8]), parse(UInt32, playerData[9]),
		parse(UInt32, playerData[10]))
		push!(playerRecords, newPlayer)
	end
end

# Output a summary line indicating how many players were found in the file
println(length(playerRecords), " PLAYERS FOUND IN FILE")

# Print a report by highest OPS and one by highest batting average
printAllPlayerStats(sort(playerRecords, by=calcOPS, rev=true), "highest OPS")
printAllPlayerStats(sort(playerRecords, by=calcBattingAvg, rev=true), "highest batting average")
println("\nEnd of Program 2")