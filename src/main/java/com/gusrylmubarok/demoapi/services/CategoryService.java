package com.gusrylmubarok.demoapi.services;

import java.util.Optional;

import com.gusrylmubarok.demoapi.models.entity.Category;
import com.gusrylmubarok.demoapi.models.repository.CategoryRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class CategoryService {

  @Autowired
  private CategoryRepository categoryRepository;
  
  public Category save(Category category) {
    return categoryRepository.save(category);
  }

  public Category findOne(Long id) {
    Optional<Category> category = categoryRepository.findById(id);
    if(!category.isPresent()){
      return null;
    }

    return category.get();
  }

  public Iterable<Category> findAll() {
    return categoryRepository.findAll();
  }

  public void removeOne(Long id) {
    categoryRepository.deleteById(id);
  }

  
}
