The Challenger

**NOTE**  
In this test I dont sure about rules no.5 that said "It IS possible for repeats to overlap" <br>
From your in this example case <br>"GAGGCATCATCATCATCATTGCC" <br>
and your result is
 {'4-CAT': 5} <br>
Then I think if it can overlap the result should have more than this because we have another tamdem repeat with "ATC" start with index 5 and TCA start with index 6

**HOW TO RUN IT:**  
1. Install go lang if you dont have
2. cd in into Harder folder 
3. Run command and input your text that you want to find temdem repeat.<br>
    ```
    "go run main.go" 
    ```
4. If you wan a binary file you can run command.<br>
    ```
    "go build" 
    ```
<br>

**HOW TO RUN TESTS**
1. cd in into Harder folder 
2. Run command.<br>
    ```
    "go test" 
    ```

**Choices**
1. I create frame of tandem word with size from 3 to 10 and loop along in text with taht frame and add start index for each word in to map that have tandem word as key and have list of start index as value.
2. Loop in map and check tandem repeat by start index list and size of that tandem word<br>
    Example map{"ACT" : [3,6,9,17,20] <br>
    from this I will know tandem word "ACT" have two repeat point in string. <br>First point start index is 3 repeat times is 3
    <br>Second point start index is 17 repeat times is 2

<br>

**LIMITATIONS**

    Range of tandem word size it can be between 3 and 10. 

**Algorithm performance**

Have two main loops with O(10N)

    O(10N) + O(10N) -> O(N)
