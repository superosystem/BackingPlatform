package com.gusrylmubarok.demoapi.models.repository;

import com.gusrylmubarok.demoapi.models.entity.Category;

import org.springframework.data.repository.CrudRepository;

public interface CategoryRepository extends CrudRepository<Category, Long>{
  
}
