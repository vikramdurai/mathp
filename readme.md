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

    Usage: `GET /api/<grade>/<syllabus>/<mode>/<pattern>/<amount>`. For example:
    `GET /api/3/NCERT/Algebra/polynomial/2` gets two Algebra polynomial problems
    that concern the NCERT syllabus for the 3rd grade.
    
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
            {
                "question": "2ca²b + 4b²ca² = 24.3",
                "answer": 3.53,
            },
            {
                "question": "7x - 3y2z + 3z = 43",
                "answer": 35,
            }
        ]
    }
    ```