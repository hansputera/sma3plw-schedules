import os
from sys import exit as exitProgram

# TODO: break changes for clock
basePath = os.path.join(os.path.dirname(os.path.abspath(__file__)), "../")

kelas_input = input("Folder kelas: ").lower()
hari_input = input("Nama hari: ").lower()

kelas_path = os.path.join(basePath, kelas_input)
hari_path = os.path.join(basePath, kelas_input, hari_input)

checkSmth = lambda: os.path.exists(kelas_path) and os.path.exists(hari_path) == False
def get_clock(index, day):
    day = day.replace(".txt", "")
    if day.lower() == "senin":
        return [
            "08.00 - 08.40",
            "08.40 - 09.20",
            "09.20 - 10.00",
            "10.00 - 10.30",
            "10.30 - 11.10",
            "11.10 - 11.50",
            "11.50 - 12.45",
            "12.45 - 13.25",
            "13.25 - 14.05",
            "14.05 - 14.55"
        ][index]
    elif day.lower() != "jumat":
        return [
            "07.15 - 07.50",
            "07.50 - 08.25",
            "08.25 - 09.00",
            "09.00 - 09.35",
            "10.00 - 10.35",
            "10.35 - 11.10",
            "11.10 - 11.45",
            "12.45 - 13.20",
            "13.20 - 13.55",
            "13.55 - 14.30"
        ][index]
    else:
        return [
            "08.30 - 09.05",
            "09.05 - 09.40",
            "09.40 - 10.15",
            "10.35 - 11.10",
            "11.10 - 10.45"
        ][index]

if checkSmth() == False:
    print("Kelas atau Hari Invalid!")
    exitProgram(0)

i=1
file = open(hari_path, mode="w+")
wouldBreak = True

while wouldBreak:
    try:
        waktu = get_clock(i-1, os.path.basename(hari_path))
        i += 1
        content = input("")
        if content.lower() == "stop":
            wouldBreak = False
            break
        file.write(f"{waktu} | {content}\n")
    except:
        print("maximum reach, breaking")
        wouldBreak=False
        break

file.close()
