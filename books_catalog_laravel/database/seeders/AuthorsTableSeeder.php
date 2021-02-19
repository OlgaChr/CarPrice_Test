<?php

namespace Database\Seeders;

use App\Models\Author;
use Illuminate\Database\Seeder;

class AuthorsTableSeeder extends Seeder
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

        $faker = \Faker\Factory::create();

        for ($i = 0; $i < 5; $i++) {
            Author::create([
                'surname' => $faker->lastName,
                'name' => $faker->firstName,
                'birth_year' => $faker->year,
            ]);
        }
    }
}
