/*********************************************************************************************************************
main algorithm depends on optimization
1) i didnt want to take all data to memory for searching sentence
2) i take chunk to chunk to memory as a lenght of data equals to sentence length(search pattern)
3) algoritm depends on creating new chunks from previous and next chunks
  for example lets say
  text    : i want to go to school but it think i am ill
  pattern : i am ill
  
  datas will be like these
  next sentence = want to go
  next sentence = to go to
  next sentence = go to school
  
  then we will find "i am ill"
  
4) for weight of pattern i am take an algorithm functions 

  for example lets say
  text    = collection of well-known quotations that are associated with the well-known quotations is that associated
  pattern = well-known quotations is associated
  result = %50 in first part, %75 in second part and then result is %75
  
  for example lets say
  text    = "the well-known quotations is that associated for insane collection of well-known quotations that are"
  pattern = "there well-known quotations is"
  result  = "well-known quotations is" => %75 not %50
 
*****************************************************************************************************************************/

const fs = require('fs');
const stream = require('stream');

const folderPath = process.argv[2]
const searchPattern = process.argv[3]

fs.readdir(folderPath,(err, files) => {
  if(err){
    console.log("we got error",err)
    return
  } 
  files.forEach(file => {
     startSearchEngine(folderPath,file,searchPattern)
  });
});

var startSearchEngine = (folderPath,file,searchPattern)  => {
  let hashOfPattern = getSearchPatternHash(searchPattern)
  let lengthOfPattern = Buffer.from(searchPattern).length
  console.log(hashOfPattern)
 
  let previousChunk = ""
  const readabaleContents = fs.createReadStream(folderPath+"/"+file,{ highWaterMark: lengthOfPattern*2});
  readabaleContents.on('data', (chunk)=> {

    if(chunk.toString().split(' ').includes(searchPattern.split(" ")[0])){
      chunk.toString().indexOf(searchPattern.split(" ")[0])


    }

    let hashChunk = getSearchPatternHash(chunk.toString())

    if(hashOfPattern === hashChunk){
      console.log("we find it")
      return
    }
    else{
      previousChunk = chunk

    }
  });
}

var s = "Barry Manilow may claim to write the songs, but it was "    
var s1 ="William Shakespeare who coined the phrases - he contributed more "
var s2 ="phrases and sayings  to the language than any other "  
var s3 ="individual, and  sayings to  the English language in daily use. Here's a " 
var s4 ="collection of well-known quotations that are associated with " 
var s5 ="Shakespeare. Most of these were the Bard's own work but he wasn't "
var s6 ="the well-known quotations is that associated for insane "


s = s+s1+s2+s3+s4+s5+s6
var pattern = "well-known quotations is associated"
let chunk=""
let temp= ""
let slicePattern=0
let we_find_it = false
let moduloResultArr = []  


/***************************************************************
 * function : maximumIncrementalSequence
 * in a given array, it finds the maximum sequential array
 * for example: [1,2,6,7,8] => 3
 * for example:[1,2,3,6,7] => 3
 * *************************************************************/
maximumIncrementalSequence = arg => {
  let incrementalCount = 0
  let firstIndex = arg[0]
  let max =[]
  for(let i=1;i<arg.length;i++){
    if(arg[i] - firstIndex === 1){
      incrementalCount++     
      firstIndex = arg[i]
    }
    else{
      max.push(incrementalCount)
      incementalCount=0
      firstIndex = arg[i]
    }
  }
  return Math.max(...max)
}


/***************************************************************
 * function : parseMultiSpace
 * multispace can prevent to find the right match 
 * *************************************************************/
parseMultiSpace = str => {
  let s = str.trim().split(' ')
  return s.reduce((prev,next) => {
      if(next === "")
         return prev
      else return prev.concat(" ",next)
  })
}


/***************************************************************
 * function : isFilter
 * find the weight of percentage of sentence occurs in file
 * *************************************************************/
isFilter = (source,destination)=> {
   let result = destination.split(' ').map(element => source.split(' ').indexOf(element))
   let incementalCount = maximumIncrementalSequence(result)
   return (incementalCount+1) / result.length
}


/***************************************************************
 * function : isFilterResult
 * source: text pattern
 * destination:sentence which will be find
 * if destination is all occurs in text it will return true
 * *************************************************************/
isFilterResult = (source,destination)=> {
   let exist = 0
   let _destination = destination.split(' ')
   let _source = source.split(' ')
   let result = _destination.map(element => _source.indexOf(element))
   if(result.includes(-1)) return false
   else{
     for(let i=0,j=1;i<result.length-1,j<result.length;i++,j++){
        exist = Math.abs(result[j]-result[i]) + exist
    }
    if(exist === result.length-1) return true
    else return false
   }
}


/***************************************************************
 * function : isFilterResult
 * source: text pattern
 * destination:sentence which will be find
 * if destination is all occurs in text it will return true
 * *************************************************************/
pattern = parseMultiSpace(pattern)
for(let i=0;i<s.length;i=i+pattern.length){
   chunk = s.substr(i,pattern.length)
   if(chunk !== pattern){
     temp = temp.concat(chunk)
     if(temp.length != pattern.length){
       let tempPattern = temp.split(' ').slice(slicePattern)
       moduloResultArr= []
       for(j=1;j<tempPattern.length;j++){
         let newIteratifTemp = tempPattern.slice(j).join(' ')
         if(newIteratifTemp.length < pattern.length)break
         if(isFilterResult(newIteratifTemp,pattern)){
            console.log("find it with %100") 
            we_find_it =true
            break
         }
         else{
           moduloResultArr.push(isFilter(newIteratifTemp,pattern))
         }
       }
       slicePattern++
       if(we_find_it) break
     }
   }
   else {
      console.log("we find it with %100",pattern) 
      we_find_it =true
      break
   }
}
if(!we_find_it){
  console.log("we find it with %",100*Math.max(...moduloResultArr))
}
