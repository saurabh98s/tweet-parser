hr@wednesday.is


### Example cURL request
```
curl --location --request PUT '127.0.0.1:8000/split_tweets' \
--header 'Content-Type: text/plain' \
--data-raw 'After thoroughly study of humanitarian intervention over analysis of responsibility to protect the come to an end where all this argument leads to doctrine of responsibility to protect has not succeeded in establishing agreed upon conditions and methods for intervening in case of humanitarian crisis. The complete failure of United Nations peacekeeping and security council after numerous need the cries of slaughtered millions from show many years to current situation became worst slaughter in Biafra, Uganda, Burnudi, Indonesia, Bangladesh and Sudan in which current era Afghanistan, Russia Ukraine war we will finally say that legitimateness Of these mass '
```


### Sample Response
```
[
    {
        "label": 1,
        "Tweet": "After thoroughly study of humanitarian intervention"
    },
    {
        "label": 2,
        "Tweet": " over analysis of responsibility to protect the co"
    },
    {
        "label": 3,
        "Tweet": "me to an end where all this argument leads to doct"
    },
    {
        "label": 4,
        "Tweet": "rine of responsibility to protect has not succeede"
    },
    {
        "label": 5,
        "Tweet": "d in establishing agreed upon conditions and metho"
    },
    {
        "label": 6,
        "Tweet": "ds for intervening in case of humanitarian crisis."
    },
    {
        "label": 7,
        "Tweet": " The complete failure of United Nations peacekeepi"
    },
    {
        "label": 8,
        "Tweet": "ng and security council after numerous need the cr"
    },
    {
        "label": 9,
        "Tweet": "ies of slaughtered millions from show many years t"
    },
    {
        "label": 10,
        "Tweet": "o current situation became worst slaughter in Biaf"
    },
    {
        "label": 11,
        "Tweet": "ra, Uganda, Burnudi, Indonesia, Bangladesh and Sud"
    },
    {
        "label": 12,
        "Tweet": "an in which current era Afghanistan, Russia Ukrain"
    },
    {
        "label": 13,
        "Tweet": "e war we will finally say that legitimateness Of t"
    },
    {
        "label": 14,
        "Tweet": "hese mass "
    }
]


```