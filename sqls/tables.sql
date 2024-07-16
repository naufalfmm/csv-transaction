CREATE TABLE IF NOT EXISTS balances (
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    total DECIMAL(15,3) NOT NULL,
    failed_total DECIMAL(15,3) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    balance_id BIGINT NOT NULL,
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    amount DECIMAL(15,3) NOT NULL,
    status VARCHAR(255) NOT NULL,
    notes VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    
    CONSTRAINT fk_transaction_balance FOREIGN KEY (balance_id) REFERENCES balances(id)
);