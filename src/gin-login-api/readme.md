| **Method** | **Endpoint** | **Purpose**                  | **Auth Required?** |
| ---------------- | ------------------ | ---------------------------------- | ------------------------ |
| **POST**   | `/api/register`  | Register a new user                | ❌ No                    |
| **POST**   | `/api/login`     | Authenticate user & return JWT     | ❌ No                    |
| **GET**    | `/api/profile`   | Get user profile                   | ✅ Yes (JWT)             |
| **POST**   | `/api/logout`    | Invalidate user session (optional) | ✅ Yes (JWT)             |

## Stress testing

```shell
hey -n 1000 -c 50 -m GET -H "Content-Type: application/json" -d '{"email": "sumant@mail.com", "password": "123456"}' "http://localhost:8080/api/login"

```
