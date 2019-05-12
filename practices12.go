const fs = require('fs');
const crypto = require('crypto');
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

var getSearchPatternHash = searchPattern => {
  return sha256Hash(searchPattern,"pattern with length" + searchPattern.split(' ').length)
}


var sha256Hash = function (str,hashKey){  
  var hashedStr = crypto.createHmac('sha256',str).update(hashKey).digest('hex');
  return hashedStr;
}


var s = "Barry Manilow may claim to write the songs, but it was "    
var s1 ="William Shakespeare who coined the phrases - he contributed more "
var s2 ="phrases and sayings to the English language than any other "  
var s3 ="individual, and most of them are still in daily use. Here's a " 
var s4 ="collection of well-known quotations that are associated with " 
var s5 ="Shakespeare. Most of these were the Bard's own work but he wasn't"

s = s+s1+s2+s3+s4+s5
var pattern = "Shakespeare who phrases"
var l = pattern.length
let chunk=""
let temp= ""
let j = 1
let we_find_it = false
let moduloResultArr = []  

findIncremental = arg => {
  let incementalCount = 0
  let firstIndex = arg[0]
  for(let i=1;i<arg.length;i++){
       if(arg[i] - firstIndex === 1)
         incementalCount++
       firstIndex = arg[i]
  }
  return incementalCount
}

isFilter = (source,destination)=> {
   let exist = 0
   let _destination = destination.split(' ')
   let _source = source.split(' ')
   let result = _destination.map(element => _source.indexOf(element))
   let len = _destination.length
   incementalCount = findIncremental(result)
   return (incementalCount+1) / result.length
}

isFilterResult = (source,destination)=> {
   let exist = 0
   let _destination = destination.split(' ')
   let _source = source.split(' ')
   let result = _destination.map(element => _source.indexOf(element))
   if(result.includes(-1)) return false
   else{
    let len = _destination.length
    for(let i=0;i<len;i++){
        exist = result[len-1]-result[len-2]
    }
    if(exist === 1) return true
    else return false
   }
}
     
for(let i=0;i<s.length;i=i+l){
   chunk = s.substr(i,l)
   if(chunk !== pattern){
     temp = temp.concat(chunk)
     if(temp.length != pattern.length){
       let tempPattern = temp.split(' ')
       for(;j<tempPattern.length;j++){
         let newIteratifTemp = tempPattern.slice(j).join(' ')
         if(newIteratifTemp.length < pattern.length)break
         moduloResultArr.push(isFilter(newIteratifTemp,pattern))
         if(isFilterResult(newIteratifTemp,pattern)){
            console.log("we find it with %100") 
            we_find_it =true
            break
         } 
       }
       if(we_find_it) break
     }
   }
   else {
      console.log("we find it with %100") 
      break
   }
}
if(!we_find_it){
  console.log("we find it with %",Math.floor(100*Math.max(...moduloResultArr)))
}
   

     
     
     
  









