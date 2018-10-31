# mathp
Math problem generation.

## API Calls and Parameters
- **/api/**:
    the API endpoint. Uses the following parameters:

    | Parameter |    Sample Values        |    Type     |
    |-----------|-------------------------|-------------|
    | Grade     |    `5`,`3`,`6`          |    `int`    |
    | Syllabus  |   `NCERT`,`CBSE`        |   `string`  |
    | Mode      |`Algebra`,`Geometry`     |   `string`  |
    | Pattern   | `polynomial`,`binomial` |   `string`  |
    | Amount    |   `3`,`4`,`7`,`9`       |   `int`     |

    Usage:
    ```
    GET /api/?grade=<grade>&syllabus=<syllabus>&mode=<mode>;pattern=<pattern>&amount=<amount>
    ```
    For example:
    ```
    GET /api/?grade=3&syllabus=ncert&mode=algebra&pattern=polynomial&amount=2
    ```
    gets two Algebra polynomial problems that concern the NCERT syllabus for the 3rd grade.
    
    A typical result:
    ```
    {
        "request": {
            "grade": 3,
            "syllabus": "NCERT",
            "mode": "Algebra",
            "pattern": "polynomial",
            "amount": 2,
        },
        "reply": [
            "3x + -4y",
            "4a² - 2bc²a",
        ]
    }
    ```