from os import path, listdir
from json import dumps
from re import sub

guru_list = {}
files = [f for f in listdir(
    path.abspath(f"{__file__}/../..")
) if f.endswith('.txt')]

for file in files:
    f = open(path.abspath(f"{__file__}/../../{file}"))
    contents = f.read()
    contents = contents.splitlines()

    for c in contents:
        c = c.split('|')
        
        ID = sub(r'\s', '', c[0])
        nama = c[1].strip()

        print(f"Adding {ID} - {nama}")
        if ID in guru_list:
            guru_list[ID]['mapel'].append(file.replace('.txt', ''))
            continue
        
        jam_mengajar = '-'
        ket = '-'

        if len(c) == 4:
            jam_mengajar = sub(r'\s', '', c[2].strip())
            ket = c[3].strip()
        else:
            ket = c[2].strip()
        
        guru_list[ID] = {
            "nama": nama,
            "mapel": [file.replace('.txt', '')],
            "jam_mengajar": int(jam_mengajar) if jam_mengajar != '-' else '-',
            "keterangan": ket
        }
    
    f.close()

wf = open(path.abspath(f"{__file__}/../../raw.json"), "w+")
wf.write(dumps(guru_list, indent=2))
wf.close()
