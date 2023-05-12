CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       login TEXT NOT NULL UNIQUE,
                       password_hash TEXT NOT NULL,
                       role TEXT NOT NULL
);

CREATE TABLE students (
                         student_id SERIAL PRIMARY KEY,
                         specialty VARCHAR NOT NULL,
                         avg_mark NUMERIC NOT NULL,
                         student_name VARCHAR NOT NULL,
                         student_email VARCHAR NOT NULL,
                         student_phone_number VARCHAR NOT NULL,
                         speciality TEXT NOT NULL,
                         CONSTRAINT student_fk0 FOREIGN KEY (student_id) REFERENCES users (id)
);

CREATE TABLE companies (
                         company_id SERIAL PRIMARY KEY,
                         name VARCHAR NOT NULL,
                         description TEXT NOT NULL,
                         address VARCHAR NOT NULL,
                         contact_person VARCHAR NOT NULL,
                         phone_number VARCHAR NOT NULL,
                         email VARCHAR NOT NULL,
                         website VARCHAR NOT NULL,
                         CONSTRAINT company_fk0 FOREIGN KEY (company_id) REFERENCES users (id)
);



CREATE TABLE requests (
                         id SERIAL PRIMARY KEY,
                         company_id INT NOT NULL,
                         student_id INT NOT NULL,
                         status VARCHAR NOT NULL,
                         projects_link TEXT,
                         language_skills TEXT,
                         work_experience INTEGER,
                         skills TEXT,
                         CONSTRAINT request_fk0 FOREIGN KEY (company_id) REFERENCES companies (company_id),
                         CONSTRAINT request_fk1 FOREIGN KEY (student_id) REFERENCES students (student_id)
);

CREATE TABLE practices (
                          id SERIAL PRIMARY KEY,
                          company_id INTEGER NOT NULL,
                          duration TEXT NOT NULL,
                          start_date DATE NOT NULL,
                          is_payment BOOLEAN NOT NULL,
                          payment_amount TEXT,
                          working_hours TEXT,
                          expectations TEXT,
                          benefits TEXT,
                          restrictions TEXT,
                          required_skills TEXT,
                          additional_information TEXT,
                          CONSTRAINT practice_fk0 FOREIGN KEY (company_id) REFERENCES companies (company_id)
);

CREATE TABLE universities (
                            university_id SERIAL PRIMARY KEY,
                            CONSTRAINT university_fk0 FOREIGN KEY (university_id) REFERENCES users (id)
);