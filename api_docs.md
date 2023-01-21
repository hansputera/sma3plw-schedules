## API Docs

1. To create schedule, just hit `/schedules/:class/:day` with POST method
Payload:
```json
[
    {
        "begin": "07.15",
        "end": "08.10",
        "code": 15
    }, {
        
    }
]
```

Just fill out the `code` with teacher's code, or set it to `0` if it's _break_ time.

2. To generate _times_, hit `/times/:day`
Response:
```json
{
    "data": [
        "07.15 - 07.45",
        "07.45 - 08.10",
        "..."
    ]
}
```

3. To get teachers data, hit `/teachers`
Response:
```json
{
    "data": {
        "81": {
            "nama": "Drs. Yobelt Onivas",
            "mapel": ["matematika"],
            "jam_mengajar": 22,
            "keterangan": "Matematika PMINATAN (X IPA 2&7) ..."
        },
        "..": {

        }
    }
}
```

4. To get all schedule, hit `/schedules/:class` (add `?teachers=1` to get the teacher's data)
Response:
```json
{
    "data": [
        {
            "urutan": 1,
            "code": 81,
            "beginAt": "07.15",
            "endAt": "08.30"
        }, {

        }
    ]
}
```

5. Push updates to git, hit `/push_updates`