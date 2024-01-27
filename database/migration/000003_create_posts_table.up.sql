CREATE TABLE IF NOT EXISTS posts (
	id BIGSERIAL PRIMARY KEY,
    uuid UUID UNIQUE DEFAULT gen_random_uuid(),
    tag_uuid UUID NOT NULL,
    user_uuid UUID NOT NULL,
    title VARCHAR(255) NOT NULL,
    thumbnail TEXT,
    content TEXT NOT NULL,
    slug TEXT NOT NULL,
    keyword TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    is_highlight BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    deleted_at TIMESTAMP WITH TIME ZONE
);