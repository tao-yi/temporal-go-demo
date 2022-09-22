CREATE DATABASE "Payment";

\connect "Payment";

CREATE TABLE "CashFlow" (
  "id" SERIAL NOT NULL,
  "from_account" INT,
  "to_account" INT,
  "amount" DECIMAL
);

CREATE TABLE "Account" (
  "id" SERIAL NOT NULL,
  "balance" DECIMAL
);

INSERT INTO "Account" ("balance")
VALUES (100), (100);