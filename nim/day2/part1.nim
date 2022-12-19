import os

type Move {.pure} = enum
  ROCK = 1, PAPER = 2, SCISSORS = 3

proc processInput(filename:string):void
proc processMove(playerMove:Move, opponentMove:Move):int
proc moveFromChar(c:char):Move

proc main():void =
  echo "AoC Day 2 Part 1"
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
    score += processMove(line[2].moveFromChar,line[0].moveFromChar)

  echo score


proc moveFromChar(c:char):Move =
  if c == 'A' or c == 'X': return ROCK
  if c == 'B' or c == 'Y': return PAPER
  if c == 'C' or c == 'Z': return SCISSORS

proc processMove(playerMove:Move, opponentMove:Move):int = 
  if playerMove == opponentMove: return ord(playerMove) + 3

  case playerMove:
    of ROCK:
      if opponentMove == PAPER: return ord(playerMove)
    of PAPER:
      if opponentMove == SCISSORS: return ord(playerMove)
    of SCISSORS:
      if opponentMove == ROCK: return ord(playerMove)
  
    

  return ord(playerMove) + 6
  
main()