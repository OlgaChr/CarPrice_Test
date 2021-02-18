<?php

namespace Database\Seeders;

use App\Models\Author;
use App\Models\Book;
use Illuminate\Database\Seeder;

class BooksTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        // Удалим имеющиеся в таблице данные
        Author::truncate();
        Book::truncate();

        $faker = \Faker\Factory::create();

        for ($i = 0; $i < 5; $i++) {
            Author::create([
                'surname' => $faker->lastName,
                'name' => $faker->firstName,
                'birth_year' => $faker->year,
            ]);
        }

        for ($i = 0; $i < 10; $i++) {
            Book::create([
                'name' => $faker->sentence,
                'publication_year' => $faker->year,
                'summary' => $faker->paragraph,
                'author_id' => $faker->randomElement([1, 2, 3, 4, 5]),
            ]);
        }
    }
}
