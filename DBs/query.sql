-- ============================================
-- =============== UBICACIONES =================
-- ============================================

CREATE TABLE countries
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE departments
(
    id         SERIAL PRIMARY KEY,
    country_id INTEGER      NOT NULL REFERENCES countries (id),
    name       VARCHAR(100) NOT NULL,
    UNIQUE (country_id, name)
);

CREATE TABLE municipalities
(
    id            SERIAL PRIMARY KEY, ,
    department_id INTEGER      NOT NULL REFERENCES departments (id),
    name          VARCHAR(100) NOT NULL,
    UNIQUE (department_id, name)
);

CREATE TABLE neighborhoods
(
    id              SERIAL PRIMARY KEY, ,
    municipality_id INTEGER      NOT NULL REFERENCES municipalities (id),
    name            VARCHAR(100) NOT NULL,
    UNIQUE (municipality_id, name)
);

-- ============================================
-- =============== USUARIOS ====================
-- ============================================

CREATE TABLE users
(
    id                    BIGSERIAL PRIMARY KEY,
    user_name             TEXT,
    first_name            TEXT,
    last_name             TEXT,
    email                 TEXT,
    country_code          TEXT,
    phone_number          TEXT,
    password_hash         TEXT,
    user_type             TEXT,
    profile_picture_url   TEXT,
    bio                   TEXT,
    birth_date            TIMESTAMP WITH TIME ZONE,
    gender                TEXT,
    specialties           JSONB,
    account_type          VARCHAR(50),
    tasks_completed       BIGINT  DEFAULT 0,
    theme_preference      VARCHAR(20),
    verification_code     VARCHAR(6),
    verification_attempts INTEGER DEFAULT 0,
    code_expiration       VARCHAR(15)  NOT NULL,
    message_id            VARCHAR(255) NOT NULL,
    is_active             BOOLEAN DEFAULT FALSE,
    is_verified           BOOLEAN DEFAULT FALSE,
    is_blocked            BOOLEAN DEFAULT FALSE,
    last_login_at         TIMESTAMP(0),
    created_at            TIMESTAMP(0),
    updated_at            TIMESTAMP(0)
);

CREATE TABLE roles
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL UNIQUE,
    level       INTEGER     NOT NULL CHECK (level > 0),
    description TEXT,
    created_at  TIMESTAMP(0),
    updated_at  TIMESTAMP(0)
);

-- ============================================
-- =============== DIRECCIONES =================
-- ============================================

CREATE TABLE user_addresses
(
    id                      SERIAL PRIMARY KEY,
    user_id                 INTEGER      NOT NULL REFERENCES users (id),
    label                   VARCHAR(100) NOT NULL,
    street_number           VARCHAR(50),
    apartment_details       VARCHAR(100) NOT NULL,
    additional_instructions VARCHAR(150),
    country_id              INTEGER      NOT NULL REFERENCES countries (id),
    department_id           INTEGER      NOT NULL REFERENCES departments (id),
    municipality_id         INTEGER      NOT NULL REFERENCES municipalities (id),
    neighborhood_id         INTEGER      NOT NULL REFERENCES neighborhoods (id),
    is_primary              BOOLEAN   DEFAULT FALSE,
    created_at              TIMESTAMP DEFAULT NOW(),
    updated_at              TIMESTAMP
);

-- ============================================
-- =============== TASK REQUESTS ===============
-- ============================================

CREATE TABLE task_requests
(
    id                 SERIAL PRIMARY KEY,
    user_id            INTEGER      NOT NULL REFERENCES users (id),
    title              VARCHAR(100) NOT NULL,
    description        TEXT         NOT NULL,
    service_date       DATE         NOT NULL,
    service_time_range VARCHAR(50)  NOT NULL,
    address_id         INTEGER      NOT NULL REFERENCES user_addresses (id),
    location_text      VARCHAR(255),
    preferences        JSONB                 DEFAULT '{}'::jsonb,
    attached_files     JSONB                 DEFAULT '[]'::jsonb,
    budget_proposal    NUMERIC(10, 2)
        CHECK (budget_proposal >= 0),
    status             VARCHAR(20)  NOT NULL DEFAULT 'pendiente',
    file_url           TEXT         NOT NULL,
    file_type          VARCHAR(20)
        CHECK (file_type IN ('image', 'pdf')),
    created_at         TIMESTAMP             DEFAULT NOW(),
    updated_at         TIMESTAMP             DEFAULT NOW()
);

-- ============================================
-- =============== REVIEWS =====================
-- ============================================

CREATE TABLE reviews
(
    id                   SERIAL PRIMARY KEY,
    reviewer_id          INTEGER      NOT NULL REFERENCES users (id),
    reviewed_user_id     INTEGER      NOT NULL REFERENCES users (id),
    task_request_id      INTEGER      NOT NULL REFERENCES task_requests (id) ON DELETE CASCADE,
    comment              VARCHAR(100) NOT NULL,
    overall_rating       NUMERIC(2, 1) CHECK (overall_rating >= 0 AND overall_rating <= 5),
    communication_rating NUMERIC(2, 1) CHECK (communication_rating >= 0 AND communication_rating <= 5),
    punctuality_rating   NUMERIC(2, 1) CHECK (punctuality_rating >= 0 AND punctuality_rating <= 5),
    respect_rating       NUMERIC(2, 1) CHECK (respect_rating >= 0 AND respect_rating <= 5),
    quality_rating       NUMERIC(2, 1) CHECK (quality_rating >= 0 AND quality_rating <= 5),
    equipment_rating     NUMERIC(2, 1) CHECK (equipment_rating >= 0 AND equipment_rating <= 5),
    created_at           TIMESTAMP DEFAULT NOW(),
    updated_at           TIMESTAMP DEFAULT NOW(),
    UNIQUE (reviewer_id, task_request_id)
);

-- ============================================
-- =============== CATEGORÍAS ==================
-- ============================================

CREATE TABLE category
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(75)            NOT NULL,
    description VARCHAR(255)           NOT NULL,
    is_active   BOOLEAN   DEFAULT TRUE NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP
);

CREATE TABLE subcategory
(
    id          SERIAL PRIMARY KEY,
    category_id INTEGER
) NOT NULL REFERENCES category (id),
    name        VARCHAR(75) NOT NULL,
    description VARCHAR(255) NOT NULL,
    image_url   TEXT      ) NOT NULL,
    is_active   BOOLEAN   DEFAULT TRUE NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP
);

-- ============================================
-- ============= UI / MENÚ PERFIL =============
-- ============================================

CREATE TABLE profile_menu_item
(
    id          SERIAL PRIMARY KEY,
    screen_code VARCHAR(50)
) NOT NULL,
    element_code VARCHAR(50)) NOT NULL,
    user_type    VARCHAR(20)) NOT NULL
        CHECK (user_type IN ('client', 'tasker', 'all')),
    group_label  VARCHAR(50)) NOT NULL,
    icon_code    VARCHAR(50)) NOT NULL,
    order_index  INTEGER) NOT NULL CHECK (order_index >= 0),
    is_active    BOOLEAN   DEFAULT TRUE  NOT NULL,
    created_at   TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at   TIMESTAMP,
    UNIQUE (screen_code, element_code, user_type)
);

CREATE TABLE ui_text
(
    id           SERIAL PRIMARY KEY,
    screen_code  VARCHAR(50) NOT NULL,
    element_code VARCHAR(50)
) NOT NULL,
    content      VARCHAR(150) NOT NULL,
    user_type    VARCHAR(20)) NOT NULL CHECK (user_type IN ('client', 'tasker', 'all')),
    language     VARCHAR(10)) NOT NULL DEFAULT 'es',
    order_index  INTEGER    ) NOT NULL CHECK (order_index >= 0),
    is_active    BOOLEAN   DEFAULT TRUE  NOT NULL,
    created_at   TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at   TIMESTAMP,
    UNIQUE (screen_code, element_code, language)
);

-- ============================================
-- =========== NOTIFICACIONES =================
-- ============================================

CREATE TABLE user_notifications
(
    id      SERIAL PRIMARY KEY,
    user_id INTEGER
) NOT NULL REFERENCES users (id),
    title             VARCHAR(100) NOT NULL,
    description       TEXT       ) NOT NULL,
    image_url         TEXT,
    action_url        TEXT,
    notification_type VARCHAR(20)) NOT NULL,
    is_read           BOOLEAN   DEFAULT FALSE NOT NULL,
    is_seen           BOOLEAN   DEFAULT FALSE NOT NULL,
    read_at           TIMESTAMP,
    created_at        TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at        TIMESTAMP
);

-- ============================================
-- =============== PAGOS =======================
-- ============================================

CREATE TABLE user_payment_methods
(
    id      SERIAL PRIMARY KEY,
    user_id INTEGER
) NOT NULL REFERENCES users (id),
    payment_type VARCHAR(20)) NOT NULL, -- card, nequi, pse, cash
    reference_id INTEGER    ) NOT NULL,
    is_primary   BOOLEAN   DEFAULT FALSE NOT NULL,
    is_active    BOOLEAN   DEFAULT TRUE  NOT NULL,
    created_at   TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at   TIMESTAMP
);

CREATE TABLE user_cards
(
    id         SERIAL PRIMARY KEY,
    card_last4 VARCHAR(4)
) NOT NULL,
    card_holder_name VARCHAR(100) NOT NULL,
    expiration_date  DATE       ) NOT NULL,
    token            TEXT       ) NOT NULL,
    created_at       TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at       TIMESTAMP
);

CREATE TABLE user_nequi_accounts
(
    id           SERIAL PRIMARY KEY, ,
    phone_number VARCHAR(20)
) NOT NULL,
    email           VARCHAR(100) NOT NULL,
    document_type   VARCHAR(20)) NOT NULL,
    document_number VARCHAR(30)) NOT NULL,
    created_at      TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at      TIMESTAMP
);

CREATE TABLE user_pse_accounts
(
    id         SERIAL PRIMARY KEY,
    email      VARCHAR(100)            NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE user_cash_preferences
(
    id    SERIAL PRIMARY KEY,
    notes TEXT
) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP
);

-- ============================================
-- =========== SEGURIDAD / TOKENS =============
-- ============================================

CREATE TABLE blacklisted_tokens
(
    id      BIGSERIAL PRIMARY KEY,
    user_id BIGINT
) NOT NULL REFERENCES users (id),
    token      TEXT        ) NOT NULL,
    token_type VARCHAR(20),
    reason     VARCHAR(255),
    ip_address VARCHAR(45),
    user_agent TEXT,
    revoked_at TIMESTAMP WITH TIME ZONE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
