package com.gusrylmubarok.demoapi.controllers;

import javax.validation.Valid;

import com.gusrylmubarok.demoapi.dto.CategoryDTO;
import com.gusrylmubarok.demoapi.dto.ResponseData;
import com.gusrylmubarok.demoapi.models.entity.Category;
import com.gusrylmubarok.demoapi.services.CategoryService;

import org.modelmapper.ModelMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.Errors;
import org.springframework.validation.ObjectError;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/categories")
public class CategoryController {
  
  @Autowired
  private CategoryService categoryService;

  @Autowired
  private ModelMapper modelMapper;

  @PostMapping
  public ResponseEntity<ResponseData<Category>> create(@Valid @RequestBody CategoryDTO categoryDTO, Errors errors) {
    ResponseData<Category> responseData = new ResponseData<>();

    if(errors.hasErrors()) {
      for(ObjectError error: errors.getAllErrors()) {
        responseData.getMessages().add(error.getDefaultMessage());
      }
      responseData.setStatus(false);
      responseData.setPayload(null);

      return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(responseData);
    }
    
    Category category = modelMapper.map(categoryDTO, Category.class);
    responseData.setStatus(true);
    responseData.setPayload(categoryService.save(category));
    return ResponseEntity.ok(responseData);
  }

  @GetMapping
  public Iterable<Category> findAll() {
    return categoryService.findAll();
  }

  @GetMapping("/{id}")
  public Category findOne(@PathVariable Long id) {
    return categoryService.findOne(id);
  }

  @PutMapping
  public ResponseEntity<ResponseData<Category>> update(@Valid @RequestBody CategoryDTO categoryDTO, Errors errors) {
    
    ResponseData<Category> responseData = new ResponseData<>();
    if(errors.hasErrors()) {
      for(ObjectError error : errors.getAllErrors()) {
        responseData.getMessages().add(error.getDefaultMessage());
      }
      responseData.setStatus(false);
      responseData.setPayload(null);

      return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(responseData);
    }

    Category category = modelMapper.map(categoryDTO, Category.class);
    responseData.setStatus(true);
    responseData.setPayload(categoryService.save(category));
         
    return ResponseEntity.ok(responseData);
  }

  
}
