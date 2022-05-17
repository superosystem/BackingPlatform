package com.gusrylmubarok.spring;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.data.domain.Sort.Order;

import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

@SpringBootTest
public class UserJpaTests {

    @Autowired
    private UserRepository userRepository;

    @Test
    public void findAllUsers() {
        List<User> users = userRepository.findAll();
        assertNotNull(users);
        assertTrue(!users.isEmpty());
    }

    @Test
    public void findUserById() {
        User user = userRepository.getOne(1);
        assertNotNull(user);
    }

    @Test
    public void createUser() {
        User user = new User(null, "Budi", "budi@gmail.com", true);
        User savedUser = userRepository.save(user);
        User findUser = userRepository.findById(savedUser.getId()).get();
        assertEquals("Budi", findUser.getName());
        assertEquals("budi@gmail.com", findUser.getEmail());
    }

    @Test
    public void getUsersSortByName() {
        Sort sort = Sort.by(Sort.Direction.ASC, "name");
        List<User> users = userRepository.findAll(sort);

        assertNotNull(users);
    }

    @Test
    public void getUsersSortByNameAscAndIdDesc() {
        Order order1 = new Order(Sort.Direction.ASC, "name");
        Order order2 = new Order(Sort.Direction.DESC, "id");
        Sort sort = Sort.by(order1, order2);
        List<User> users = userRepository.findAll(sort);
        assertNotNull(users);
    }

    @Test
    public void getUserByPage() {
        Sort sort = Sort.by(Sort.Direction.ASC, "name");
        int size = 25;
        int page = 0; //zero-based page index.
        Pageable pageable = PageRequest.of(page, size, sort);
        Page<User> usersPage = userRepository.findAll(pageable);
        System.out.println(usersPage.getTotalElements()); //Returns the total amount of elements.
        System.out.println(usersPage.getTotalPages());//Returns the number of total pages.
        System.out.println(usersPage.hasNext());
        System.out.println(usersPage.hasPrevious());
        List<User> usersList = usersPage.getContent();
        assertNotNull(usersList);
    }

}
