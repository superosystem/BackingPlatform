package com.gusrylmubarok.demoapi.models.repository;

import java.util.List;

import com.gusrylmubarok.demoapi.models.entity.Product;

import org.springframework.data.repository.CrudRepository;

public interface ProductRepository extends CrudRepository<Product, Long> {

    List<Product> findByNameContains(String name);
    
}
