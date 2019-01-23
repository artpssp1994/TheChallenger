console.log("Easier Challenge")

console.log("Example 1")
data = "Be 567 yourself;79 everyone 90 else is already 96789 taken."
answer = longestWord(data)
console.log("Longest : " + answer)
answer = sumNumberInString(data)
console.log("Sum int : " + answer + "\n")

console.log("Example 2")
data = "There are 2 hard problems in computer 798 science: cache invalidation, 23422naming things, and 22 off-by-1 errors."
answer = longestWord(data)
console.log("Longest : " + answer)
answer = sumNumberInString(data)
console.log("Sum int : " + answer + "\n")

console.log("Example 3")
data = "   Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Massa placerat duis ultricies lacus. Sit amet commodo nulla facilisi nullam. Tellus elementum sagittis vitae et. Lacus sed viverra tellus in hac habitasse 57   platea dictumst. Odio  234   aenean sed adipiscing diam donec adipiscing. Feugiat pretium nibh ipsum consequat. Nunc mattis enim ut tellus elementum 343 sagittis. Habitant morbi tristique senectus et netus et malesuada fames ac. Aliquam purus sit amet luctus venenatis. Sed nisi lacus sed viverra tellus in  hac. Nulla pharetra diam sit amet nisl suscipit adipiscing bibendum est. Congue nisi vitae suscipit tellus mauris a diam maecenas. Pharetra convallis posuere morbi leo urna molestie. Amet nisl purus in mollis. Lorem mollis aliquam ut porttitor leo.\
\
Convallis a cras semper auctor. Quis eleifend quam adipiscing vitae proin sagittis nisl. Enim praesent elementum facilisis leo vel fringilla est. Tempus quam pellentesque nec nam aliquam sem et tortor. Elit pellentesque habitant morbi tristique senectus et. Nullam vehicula ipsum a arcu cursus. Faucibus vitae aliquet nec ullamcorper sit amet risus. Nulla facilisi etiam dignissim diam. Morbi tristique senectus   et netus.     At tempor commodo ullamcorper a lacus. In pellentesque massa placerat duis. Tristique et egestas quis ipsum suspendisse ultrices gravida dictum fusce.    Maecenas accumsan lacus vel facilisis volutpat est velit egestas. Amet facilisis magna etiam tempor orci. Convallis convallis tellus id interdum velit laoreet id donec. Eu turpis egestas pretium aenean. Laoreet non curabitur gravida arcu ac tortor dignissim convallis aenean.    Dui accumsan sit amet nulla facilisi morbi.\
\
Viverra suspendisse potenti nullam ac tortor vitae purus faucibus ornare. Nunc vel risus commodo viverra maecenas accumsan lacus vel facilisis. Nibh cras 21234123412341112   pulvinar mattis nunc sed blandit libero. Condimentum id venenatis a condimentum vitae sapien pellentesque habitant morbi. Ac feugiat sed lectus vestibulum mattis ullamcorper velit. Nullam eget felis eget nunc lobortis mattis aliquam faucibus purus. Pulvinar etiam non quam lacus suspendisse faucibus interdum posuere lorem.   Maecenas sed enim ut sem viverra aliquet eget. Turpis egestas pretium aenean   pharetra. Posuere sollicitudin aliquam ultrices sagittis orci. Habitasse platea dictumst vestibulum rhoncus est pellentesque elit. Velit euismod in pellentesque massa. Leo in vitae turpis massa sed. Elementum eu facilisis sed odio morbi quis commodo odio aenean. Dignissim sodales ut eu sem. Donec et odio pellentesque diam. Interdum velit euismod in pellentesque massa. Nisi vitae suscipit tellus mauris. Donec et odio pellentesque diam volutpat commodo sed egestas. Non blandit massa enim nec dui.\
\
Varius duis at consectetur lorem donec. Ac turpis egestas integer eget aliquet 3452   nibh praesent tristique. Quis imperdiet massa tincidunt nunc pulvinar sapien et ligula. Sem nulla     pharetra diam sit amet nisl suscipit adipiscing. Massa tempor nec feugiat nisl pretium fusce id velit. Sed libero enim sed faucibus turpis in eu mi bibendum. Natoque penatibus et magnis dis parturient montes. Nulla facilisi morbi tempus iaculis urna id. Integer enim neque volutpat ac tincidunt. Nisl purus in mollis nunc sed id semper risus in. Eget mauris pharetra et ultrices. "
answer = longestWord(data)
console.log("Longest : " + answer)
answer = sumNumberInString(data)
console.log("Sum int : " + answer + "\n")


function longestWord(data){
    let answer = ""
    data = data.replace(/[.,\/#!$%\^&\*;:{}=\_`~()]/g," ")
    data = data.replace(/[0-9]/g, "")
    let wordList = data.split(" ");
    wordList.forEach(element => {
        if (element.length > answer.length) {
            answer = element
        }
    });
    return answer
}

function sumNumberInString(data){
    let answer = 0
    data = data.replace(/[.,\/#!$%\^&\*;:{}=\_`~()]/g," ")
    data = data.replace(/[^0-9]/g, ' ');
    let wordList = data.split(" ");
    wordList.forEach(element => {
        if (!isNaN(parseInt(element, 10))){
            answer += parseInt(element, 10)
        }
    });
    return answer
}