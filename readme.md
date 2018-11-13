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
    | Pattern   | `polynomial`,`lineq`    |   `string`  |
    | Amount    |   `3`,`4`,`7`,`9`       |   `int`     |

    Currently we only use Pattern and Amount, so feel free to omit them in the url
    for the time being.

    Usage:
    ```
    GET /api/?pattern=<pattern>&amount=<amount>
    ```
    For example:
    ```
    GET /api/?pattern=polynomial&amount=2
    ```
    gets two Algebra polynomials.
    
    A typical result:
    ```
    {
        "request": {
            "pattern": "polynomial",
            "amount": 2,
        },
        "reply": [
            "3x + -4y",
            "4a² - 2bc²a",
        ]
    }
    ```