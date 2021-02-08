**Fetch Server**
----
Fetch server that handles various data fetch from OMDB API. Running on port `9080`

* **URL**
  /
* **Method:**
  `GET`
*  **URL Params**
    search string <br />
    page string
* **Data Params**
    None
* **Success Response:**
  * **Code:** 200 <br />
* **Error Response:**
  * **Code:** 500 Internal Server Error <br />
    **Content:** `{ error : "Internal Server Error" }`
----
* **URL**
  /:id
* **Method:**
  `GET`
*  **URL Params**
    None
* **Data Params**
    None
* **Success Response:**
  * **Code:** 200 <br />
* **Error Response:**
  * **Code:** 500 Internal Server Error <br />
    **Content:** `{ error : "Internal Server Error" }`
