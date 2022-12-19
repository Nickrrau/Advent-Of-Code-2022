import os

type Move {.pure} = enum
  ROCK = 1, PAPER = 2, SCISSORS = 3

proc processInput(filename:string):void
proc processMove(outcome:char, opponentMove:Move):int
proc moveFromChar(c:char):Move

proc main():void =
  echo "AoC Day 2 Part 2"
  let params = commandLineParams()

  if params.len < 1 :
    echo "Need to provide input"
    return
  
  processInput(params[0])

proc processInput(filename:string):void =
  let file = open(filename)
  defer: file.close()

  var score = 0

  for line in file.lines():
    score += processMove(line[2],line[0].moveFromChar)

  echo score


proc moveFromChar(c:char):Move =
  if c == 'A' or c == 'X': return ROCK
  if c == 'B' or c == 'Y': return PAPER
  if c == 'C' or c == 'Z': return SCISSORS

proc processMove(outcome:char, opponentMove:Move):int = 
  if outcome == 'Y': return ord(opponentMove) + 3

  if outcome == 'X':
    case opponentMove:
      of ROCK:
        return ord(SCISSORS)
      of PAPER:
        return ord(ROCK)
      of SCISSORS:
        return ord(PAPER)
 
  case opponentMove:
    of ROCK:
      return ord(PAPER) + 6
    of PAPER:
      return ord(SCISSORS) + 6
    of SCISSORS:
      return ord(ROCK) + 6
    

  return 0
  
main()