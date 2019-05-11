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


var s = "aslında ben sensiz olmam"
var pattern = "sensiz olmam"
var l = pattern.length
let chunk=""
let temp= ""
let j = 1
let we_find_it = false
for(let i=0;i<s.length;i=i+l){
   chunk = s.substr(i,l)
   if(chunk !== pattern){
     temp = temp.concat(chunk)
     if(temp.length != pattern.length){
       let tempPattern = temp.split(' ')
       for(;j<tempPattern.length;j++){
         let newIteratifTemp = tempPattern.slice(j).join(' ')
         if(newIteratifTemp.length < pattern.length)break
         if(newIteratifTemp.includes(pattern)){
            console.log("we find it")  
            we_find_it =true
            break
         } 
       }
       if(we_find_it) break
     }
   }
   else {
      console.log("we find it") 
      break
   }
}
   

     
     
     
     
  






