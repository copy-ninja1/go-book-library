# A Book Library Server writen with the Go Programming Language (GOLANG).

## API Documentation

## **All Books**

Returns an array of all the books in the database.

- **URL**

/api/v1/book

- **Method:**

  `GET`

- **URL Params**

  **Required:**

- **Data Params**

  None

- **Success Response:**

  - **Code:** 200 <br />
    **Content:**

    ```javascript
    [
      {
        ID: 1,
        CreatedAt: "2021-09-18T22:53:18.652583692+01:00",
        UpdatedAt: "2021-09-18T22:53:18.652583692+01:00",
        DeletedAt: null,
        title: "1984",
        author: "Geroge Orwell",
        rating: 5,
      },
    ];
    ```

    OR

  - **Code:** 200 <br />
    **Content:**
    ```javascript
    [];
    ```

- **Sample Call:**

  ```curl
  curl GET 'http://localhost:3000/api/v1/book'
  ```

## **Get A single Book**

Returns json data of a single book.

- **URL**

/api/v1/book/:id

- **Method:**

  `GET`

- **URL Params**

  **Required:**

  `id=[integer]`

- **Data Params**

  None

- **Success Response:**

  - **Code:** 200 <br />
    **Content:**
    ```json
    {
      "ID": 1,
      "CreatedAt": "2021-09-18T22:53:18.652583692+01:00",
      "UpdatedAt": "2021-09-18T22:53:18.652583692+01:00",
      "DeletedAt": null,
      "title": "1984",
      "author": "Geroge Orwell",
      "rating": 5
    }
    ```

- **Error Response:**

  - **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "Book doesn't exist" }`

- **Sample Call:**

  ```curl
  curl GET 'http://localhost:3000/api/v1/book/2'
  ```

## **Delete A Book**

Returns "Book deleted successfully".

- **URL**

  /api/v1/book/:id

- **Method:**

  `DELETE`

- **URL Params**

  **Required:**

  `id=[integer]`

- **Data Params**

  None

- **Success Response:**

  - **Code:** 200 <br />
    **Content:** `Book deleted successfully`

- **Error Response:**

  - **Code:** 404 NOT FOUND <br />
    **Content:** `No book found with given ID`

  OR

  - **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`

- **Sample Call:**

  ```curl
  curl DELETE 'http://localhost:3000/api/v1/book/3'
  ```
