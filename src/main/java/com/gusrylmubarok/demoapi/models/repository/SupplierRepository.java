package com.gusrylmubarok.demoapi.models.repository;

import com.gusrylmubarok.demoapi.models.entity.Supplier;

import org.springframework.data.repository.CrudRepository;

public interface SupplierRepository extends CrudRepository<Supplier, Long>{
  
}
