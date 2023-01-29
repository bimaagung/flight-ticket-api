CREATE TABLE IF NOT EXISTS flights (
    "id" VARCHAR(16) NOT NULL,
    "category_id" VARCHAR(16) NOT NULL,
    "flight_number" VARCHAR(10) NOT NULL,
    "departure" VARCHAR(112) NOT NULL,
    "departure time" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "arrive" VARCHAR(112) NOT NULL,
    "time_arrives" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "seats" INTEGER NOT NULL,
    "price" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL,
    "updated_at" TIMESTAMP(0) WITH TIME zone NOT NULL,
    "is_delete" BOOLEAN NOT NULL
)