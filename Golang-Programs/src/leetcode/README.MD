# Leetcode Chellanges  

Here are few of those leetcode chellanges that I have solved particularly in `golang`. I have described the problems below. The codes can be accessed through the links, provided along with the respective codes themselves.  

## 1. [Two Sum](./twoSum.go) <sub>\[easy\]</sub>  
Given an array of integers `nums` and an integer `target`, return _indices of the two numbers such that they add up to `target`_.  
You may assume that each input would have **exactly one solution**, and you may not use the same element twice.  
You can return the answer in any order.  

**Example 1:**  
```  
Input: nums = [2,7,11,15], target = 9  
Output: [0,1]  
Output: Because nums[0] + nums[1] == 9, we return [0, 1].  
```  

**Example 2:**  
```  
Input: nums = [3,2,4], target = 6  
Output: [1,2]  
```  

**Example 3:**  
```  
Input: nums = [3,3], target = 6  
Output: [0,1]  
```  

**Constraints:**  
> 2 <= nums.length <= 10<sup>3</sup>  
> -10<sup>9</sup> <= nums[i] <= 10<sup>9</sup>  
> -10<sup>9</sup> <= target <= 10<sup>9</sup>  
> **Only one valid answer exists.**  

_**access the solution [here](./twoSum.go)**_  

## 1. [Longer Contiguous Segments of Ones than Zeros](./pyVersion/longestSubsequence.py) <sub>\[easy\]</sub>  
Given a binary string `s`, return `true` if the **longest** contiguous segment of `1`s is **strictly longer** than the **longest** contiguous segment of `0`s in s. Return `false` otherwise.  
* For example, in `s = "110100010"` the longest contiguous segment of 1s has length 2, and the longest contiguous segment of 0s has length 3.  
  
Note that if there are no `0`s, then the longest contiguous segment of 0s is **considered** to have length 0. The same applies if there are no `1`s.  

**Example 1:**  
```  
Input: s = "1101"  
Output: true  
Explanation:  
The longest contiguous segment of 1s has length 2: "1101"   
The longest contiguous segment of 0s has length 1: "1101"  
The segment of 1s is longer, so return true.  
```  

**Example 2:**  
```  
Input: s = "111000"  
Output: false  
Explanation:  
The longest contiguous segment of 1s has length 3: "111000"  
The longest contiguous segment of 0s has length 3: "111000"   
The segment of 1s is not longer, so return false.  
```  

**Example 3:**  
```  
Input: s = "110100010"  
Output: false  
Explanation:  
The longest contiguous segment of 1s has length 2: "110100010"  
The longest contiguous segment of 0s has length 3: "110100010"   
The segment of 1s is not longer, so return false.   
```  

**Constraints:**  
> `1 <= s.length <= 100`  
> `s[i]` is either `'0'` or `'1'`.  

_**access the solution [here](./pyVersion/longestSubsequence.py)**_  

