import os
import parseutils

proc processInput(filename:string):void =
  let file = open(filename)
  defer: file.close()

  var 
    cal = 0
    tmp = 0

  for line in file.lines():
    if line.len < 1:
      tmp = 0
      continue
    
    var res:int
    discard parseInt(line,res,0)
    tmp += res
    
    if tmp > cal: cal = tmp
  
  echo cal


proc main():void =
  echo "AoC Day 1 Part 1"
  let params = commandLineParams()

  if params.len < 1 :
    echo "Need to provide input"
    return
  
  processInput(params[0])

  
main()