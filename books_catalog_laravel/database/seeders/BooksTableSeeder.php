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
        Book::truncate();

        $faker = \Faker\Factory::create();

//        \App\Models\Book::factory()->count(10)->create();

        for ($i = 0; $i < 10; $i++) {
            Book::create([
                'name' => $faker->sentence,
                'publication_year' => $faker->year,
                'summary' => $faker->paragraph,
            ]);
        }

        foreach (Book::all() as $book) {
            $authors = Author::inRandomOrder()->take(rand(1,3))->pluck('id');
            $book->authors()->attach($authors);
        }
    }
}
