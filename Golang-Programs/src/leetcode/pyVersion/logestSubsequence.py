class Solution:
    def checkZeroOnes(self, s: str) -> bool:
        zeroList = [0]
        oneList = [0]
        
        count = 0
        zeroFlag = False
        oneFlag = False
        
        for dig in s:
            if int(dig) == 1:
                oneFlag = True
                if zeroFlag:
                    zeroList.append(count)
                    zeroFlag = False
                    count = 0
                count += 1
                
            elif int(dig) == 0:
                zeroFlag = True
                if oneFlag:
                    oneList.append(count)
                    oneFlag = False
                    count = 0
                count += 1
                
        
        
        if zeroFlag:
            zeroList.append(count)
        elif oneFlag:
            oneList.append(count)
        
        
        #print(zeroList,oneList)
        
        
        if max(oneList) > max(zeroList):
            return True
        else:
            return False
                 
            
        
