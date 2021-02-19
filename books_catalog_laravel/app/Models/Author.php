<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Author extends Model
{
    use HasFactory;

    protected $fillable = ['surname', 'name', 'middlename', 'birth_year', 'death_year'];
    protected $hidden = ['created_at', 'updated_at', 'deleted_at'];
    protected $appends = ['books_count'];

    public function books()
    {
        return $this->belongsToMany(
            Book::class,
            'authors_books',
            'author_id',
            'book_id');
    }

    public function getBooksCountAttribute(): int
    {
        return $this->books()->count();
    }
}
