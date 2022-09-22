


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
        "tweet": "1/19:After thoroughly study of"
    },
    {
        "tweet": "2/19: humanitarian intervention over analysis of"
    },
    {
        "tweet": "3/19: responsibility to protect the come to"
    },
    {
        "tweet": "4/19: an end where all this argument leads"
    },
    {
        "tweet": "5/19: to doctrine of responsibility to"
    },
    {
        "tweet": "6/19: protect has not succeeded in"
    },
    {
        "tweet": "7/19: establishing agreed upon conditions and"
    },
    {
        "tweet": "8/19: methods for intervening in case of"
    },
    {
        "tweet": "9/19: humanitarian crisis. The complete failure"
    },
    {
        "tweet": "10/19: of United Nations peacekeeping and"
    },
    {
        "tweet": "11/19: security council after numerous need"
    },
    {
        "tweet": "12/19: the cries of slaughtered millions"
    },
    {
        "tweet": "13/19: from show many years to current"
    },
    {
        "tweet": "14/19: situation became worst slaughter in"
    },
    {
        "tweet": "15/19: Biafra, Uganda, Burnudi, Indonesia,"
    },
    {
        "tweet": "16/19: Bangladesh and Sudan in which"
    },
    {
        "tweet": "17/19: current era Afghanistan, Russia Ukraine"
    },
    {
        "tweet": "18/19: war we will finally say that"
    },
    {
        "tweet": "19/19: legitimateness Of these mass "
    }
]


```