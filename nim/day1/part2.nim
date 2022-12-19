import os
import parseutils

proc push(ar:var array[3,int], val:int):void =
  if val > ar[0]:
    ar[2] = ar[1]
    ar[1] = ar[0]
    ar[0] = val
  elif val > ar[1]:
    ar[2] = ar[1]
    ar[1] = val
  elif val > ar[2]:
    ar[2] = val

proc processInput(filename:string):void =
  let file = open(filename)
  defer: file.close()

  var 
    cal = [0,0,0]
    tmp = 0

  for line in file.lines():
    if line.len < 1:
      tmp = 0
      continue
    
    var res:int
    discard parseInt(line,res,0)
    tmp += res
    
    cal.push(tmp)
  
  echo cal[0]+cal[1]+cal[2]


proc main():void =
  echo "AoC Day 1 Part 2"
  let params = commandLineParams()

  if params.len < 1 :
    echo "Need to provide input"
    return
  
  processInput(params[0])

  
main()