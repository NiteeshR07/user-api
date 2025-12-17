# Reasoning and Design Decisions

This document explains the design approach and key technical decisions taken while building the User API, aligned with standard industry evaluation criteria such as architecture, correctness, scalability, security, and maintainability.

---

## 1. Problem Understanding & Scope
The goal was to design a RESTful API to manage user data (name and date of birth) while dynamically computing age. The scope intentionally focuses on core CRUD functionality, correctness, and clean design rather than over-engineering with unnecessary features.

---

## 2. Architecture & Design
A layered architecture was adopted to ensure separation of concerns:

- **Handler Layer**: Manages HTTP request parsing, response formatting, and status codes.
- **Service Layer**: Contains business logic such as validation flow and age calculation.
- **Repository Layer**: Handles database access and SQL interactions.

This structure improves readability, simplifies debugging, and allows independent modification or testing of each layer.

---

## 3. Technology Choices
- **GoFiber** was chosen for its high performance and minimal overhead, making it suitable for REST APIs.
- **PostgreSQL** was selected for its reliability, strong data consistency, and production readiness.
- **Uber Zap** enables structured, performant logging suitable for debugging and observability.
- **go-playground/validator** ensures request validation at the API boundary, preventing invalid data from propagating into business logic.

These choices align with commonly accepted backend engineering standards.

---

## 4. Data Modeling & Correctness
- Only essential fields (`name`, `dob`) are stored in the database.
- **Age is not persisted**; it is calculated dynamically on each request to avoid stale or redundant data.
- Date handling follows a consistent `YYYY-MM-DD` format and UTC assumptions for correctness.

This ensures data normalization and long-term accuracy.

---

## 5. API Design & REST Compliance
- Standard HTTP methods (GET, POST, PUT, DELETE) are used appropriately.
- Clear and predictable endpoint naming conventions are followed.
- Proper HTTP status codes (`200`, `201`, `204`, `400`, `404`, `500`) are returned for all scenarios.

This makes the API intuitive and easy to consume.

---

## 6. Error Handling & Validation
- Input validation is performed before business logic execution.
- Errors are returned in a consistent JSON format with meaningful status codes.
- This approach improves client-side debugging and reduces ambiguity.

---

## 7. Security Considerations
- Database credentials are externalized using environment variables.
- No sensitive information is hardcoded or committed to version control.
- The design is compatible with future additions such as authentication or rate limiting.

These practices align with basic security hygiene expected in production systems.

---

## 8. Scalability & Extensibility
- The layered design allows easy extension (e.g., pagination, filtering, authentication).
- Database interactions are isolated, enabling future optimization or migration.
- Logging and health check endpoints support deployment and monitoring scenarios.

---

## 9. Documentation & Developer Experience
- A comprehensive README provides setup steps, API usage, and examples.
- cURL commands are included for easy testing.

Clear documentation improves onboarding and maintainability.

---

## 10. Conclusion
The solution prioritizes clarity, correctness, and industry-aligned best practices. It demonstrates an understanding of backend fundamentals while remaining simple, maintainable, and extensible—meeting typical company evaluation criteria for backend engineering assignments.
