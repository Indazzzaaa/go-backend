As a developer, you should be aware of several common web security vulnerabilities to protect your applications. Here are some of the most important ones:

### 1. **Injection Attacks**

* **SQL Injection (SQLi)** : Attackers inject malicious SQL queries to access, modify, or delete database data.
* **Command Injection** : Executing arbitrary system commands via input fields.
* **NoSQL Injection** : Similar to SQLi but targets NoSQL databases like MongoDB.

### 2. **Cross-Site Scripting (XSS)**

* **Stored XSS** : Malicious scripts are stored in the database and executed when users load the page.
* **Reflected XSS** : Scripts are injected via URLs or form inputs and executed in the victim’s browser.
* **DOM-based XSS** : Manipulating the webpage DOM to execute malicious scripts.

### 3. **Cross-Site Request Forgery (CSRF)**

* Tricks a user into executing unwanted actions (e.g., transferring funds) by leveraging an authenticated session.

### 4. **Broken Authentication**

* Weak password policies, credential stuffing, and improper session management lead to unauthorized access.

### 5. **Sensitive Data Exposure**

* Storing passwords in plain text, weak encryption, and insecure data transmission (e.g., missing HTTPS).

### 6. **Security Misconfiguration**

* Default credentials, excessive privileges, unnecessary services enabled, or improper error handling.

### 7. **Broken Access Control**

* Improper access controls allow unauthorized users to perform restricted actions or access sensitive data.

### 8. **Server-Side Request Forgery (SSRF)**

* An attacker tricks a server into making requests to internal or external services, potentially leaking sensitive data.

### 9. **Denial of Service (DoS) & Distributed DoS (DDoS)**

* Overloading a server with requests to make it unavailable.

### 10. **Man-in-the-Middle (MITM) Attacks**

* Intercepting communications to steal or alter data, often due to lack of HTTPS or weak encryption.

### 11. **Insecure Deserialization**

* Exploiting weak object deserialization to execute malicious code.

### 12. **Subdomain Takeover**

* When a subdomain points to a service that is no longer in use, attackers can claim it and host malicious content.

### 13. **Clickjacking**

* Tricking users into clicking something they didn’t intend to by embedding a malicious page in an iframe.

### 14. **Business Logic Vulnerabilities**

* Exploiting application workflows (e.g., bypassing payment steps in an e-commerce site).

### 15. **Supply Chain Attacks**

* Targeting third-party libraries or dependencies used by the application.
