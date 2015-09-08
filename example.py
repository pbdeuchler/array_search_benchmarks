import string, random, time

current_micro_time = lambda x: int(round(x * 1000000))

if __name__ == "__main__":
    testList = []
    for x in range(100):
        testList.append(''.join(random.choice(string.ascii_uppercase) for i in range(12)))
    searchFor = random.randint(0, 100)
    start = time.time()
    for x in testList:
        if x == testList[searchFor]:
            break
    end = time.time()
    print(current_micro_time(end - start))
